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

var categoryCollection *mongo.Collection = config.OpenCollection(config.Client, "category")

func GetCategoryByID(c *gin.Context) {
	var category models.Category

	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//create a filter to sort the data
	filter := bson.M{"_id": objectID}
	err = categoryCollection.FindOne(context.TODO(), filter).Decode(&category)
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
	var category models.Category

	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := categoryCollection.InsertOne(context.TODO(), category)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result)
}

func UpdateCategory(c *gin.Context) {}

func DeleteCategory(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	id := c.Param("id")

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
