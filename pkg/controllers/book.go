package controllers

import (
	"context"
	"elewa/pkg/config"
	"elewa/pkg/models"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var bookCollection *mongo.Collection = config.OpenCollection(config.Client, "book")

func CreateBook(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var book models.Book
	defer cancel()

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&book); validationErr != nil {
		c.JSON(400, gin.H{"error": validationErr.Error()})
		return
	}

	newBook := models.Book{
		ID:    primitive.NewObjectID(),
		Name:  book.Name,
		Price: book.Price,
	}

	//create some extra details for the user object - created_at, updated_at, ID

	book.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	book.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	book.ID = primitive.NewObjectID()
	book.Book_id = book.ID.Hex()

	result, err := bookCollection.InsertOne(ctx, newBook)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": result})
}
func GetAllBooks(c *gin.Context) {
	var books []models.Book
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := bookCollection.Find(ctx, bson.M{})
	defer cancel()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// bind books to given array of products
	for result.Next(ctx) {
		var book models.Book
		err := result.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}

	c.JSON(200, gin.H{"success": books})

}

func GetBookById(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	defer cancel()

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := bookCollection.FindOne(ctx, primitive.M{"_id": _id})

	book := models.Book{}

	err = result.Decode(&book)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": book})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var book models.Book

	// Read update model from body request
	if err := c.BindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Create filter
	filter := bson.M{"_id": objectID}

	// Prepare update model.
	update := bson.D{
		{"$set", bson.D{{"name", book.Name}, {"price", book.Price},
			{"author", bson.D{
				{"first_name", book.Author.FirstName},
				{"last_name", book.Author.LastName},
			}},
		}},
	}

	err = bookCollection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&book)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	book.ID = objectID

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)

	defer cancel()

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := bookCollection.DeleteOne(ctx, primitive.M{"_id": _id})

	if result.DeletedCount == 0 {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, gin.H{"success": "Book deleted"})
}
