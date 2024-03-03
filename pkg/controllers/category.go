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

var categoryCollection *mongo.Collection = config.OpenCollection(config.Client, "category")

func GetCategoryByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var category models.Category
	id := c.Param("category_id")

	_id, err := primitive.ObjectIDFromHex(id)
	defer cancel()

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := categoryCollection.FindOne(ctx, bson.M{"_id": _id})

	err = result.Decode(&category)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, category)
}

func GetAllCategories(c *gin.Context) {
	var categories []models.Category

	//get all data
	cursor, err := categoryCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		//create a value to decode one document
		var category models.Category
		err := cursor.Decode(&category)
		if err != nil {
			log.Fatal(err)
		}
		//add caegory to array
		categories = append(categories, category)
	}
	c.JSON(200, categories)
}

func CreateCategory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var category models.Category
	defer cancel()

	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//create some extra details for the user object  ID
	category.ID = primitive.NewObjectID()
	category.CategoryId = category.ID.Hex()

	result, err := categoryCollection.InsertOne(ctx, category)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

func UpdateCategory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var category models.Category

	id := c.Param("category_id")

	objectID, err := primitive.ObjectIDFromHex(id)
	defer cancel()

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var UpdateCategory primitive.D

	if category.Name != "" {
		UpdateCategory = append(UpdateCategory, bson.E{Key: "name", Value: category.Name})
	}

	upsert := true
	filter := bson.M{"_id": objectID}

	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	result, err := categoryCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{Key: "$set", Value: UpdateCategory},
		},
		&opt,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

func DeleteCategory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	id := c.Param("category_id")

	objectID, err := primitive.ObjectIDFromHex(id)
	defer cancel()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//create a filter to sort the data
	filter := bson.M{"_id": objectID}
	result, err := categoryCollection.DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}
