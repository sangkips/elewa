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
	"go.mongodb.org/mongo-driver/mongo/options"
)

var bookCollection *mongo.Collection = config.OpenCollection(config.Client, "book")

func CreateBook(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var book models.Book
	defer cancel()

	if err := c.BindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//create some extra details for the user object  ID
	book.ID = primitive.NewObjectID()
	book.Book_id = book.ID.Hex()

	result, err := bookCollection.InsertOne(ctx, book)
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

	id := c.Param("book_id")
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("book_id")
	objectID, err := primitive.ObjectIDFromHex(id)
	var book models.Book
	var category models.Category
	defer cancel()

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Read update model from body request
	if err := c.BindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var UpdateBook primitive.D

	if book.Name != "" {
		UpdateBook = append(UpdateBook, bson.E{Key: "name", Value: book.Name})
	}

	if book.Price != nil {
		UpdateBook = append(UpdateBook, bson.E{Key: "price", Value: book.Price})
	}

	if book.Author != nil {
		UpdateBook = append(UpdateBook, bson.E{Key: "author", Value: book.Author})
	}
	if book.CategoryId != nil {
		err := categoryCollection.FindOne(ctx, bson.M{"categoryId": book.CategoryId}).Decode(&category)
		defer cancel()

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	upsert := true
	filter := bson.M{"_id": objectID}

	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	result, err := bookCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{Key: "$set", Value: UpdateBook},
		},
		&opt,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

func DeleteBook(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("book_id")
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
