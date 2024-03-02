package controllers

import (
	"context"
	"elewa/pkg/config"
	"elewa/pkg/helper"
	"elewa/pkg/models"
	"elewa/pkg/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.OpenCollection(config.Client, "user")
var validate = validator.New()

func GetUsers(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	result, err := userCollection.Find(ctx, bson.M{})
	defer cancel()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var allUsers []bson.M
	if err = result.All(ctx, &allUsers); err != nil {
		log.Fatal(err)
	}
	c.JSON(200, allUsers)

}

func GetUser(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	userId := c.Param("user_id")

	var user models.User

	err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)

	defer cancel()
	if err != nil {
		c.JSON(500, gin.H{"error": "error while listing user"})
	}
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("user_id")
	var user models.User
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(userId)

	//validate request
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if validationErr := validate.Struct(&user); validationErr != nil {
		c.JSON(400, gin.H{"error": validationErr.Error()})
		return
	}

	update := bson.M{"first_name": user.FirstName, "last_name": user.LastName, "phone_number": user.PhoneNumber}

	result, err := userCollection.UpdateOne(ctx, bson.M{"_id": _id}, bson.M{"$set": update})

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var updatedUser models.User
	if result.MatchedCount == 1 {
		err = userCollection.FindOne(ctx, bson.M{"_id": _id}).Decode(&updatedUser)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, gin.H{"success": updatedUser})

}

func DeleteUser(c *gin.Context) {}

func RegisterUser(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User
	defer cancel()

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(400, gin.H{"error": validationErr.Error()})
		return
	}

	count1, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	defer cancel()
	if err != nil {
		log.Panic(err)
		c.JSON(500, gin.H{"error": "error while checking the email"})
		return
	}
	//hash password
	password, err := utils.GenerateHashPassword(*user.Password)
	user.Password = &password

	if err != nil {
		log.Fatal(err)
	}

	//check if the phoneNumber has already been used by another user

	count, err := userCollection.CountDocuments(ctx, bson.M{"phone": user.PhoneNumber})
	defer cancel()
	if err != nil {
		log.Panic(err)
		c.JSON(500, gin.H{"error": "phoneNumber already used"})
		return
	}

	if count1 > 0 {
		c.JSON(500, gin.H{"error": "Email already exsits"})
		return
	}
	if count > 0 {
		c.JSON(500, gin.H{"error": "Email already exsits"})
		return
	}

	//create some extra details for the user object - created_at, updated_at, ID

	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.UserID = user.ID.Hex()

	//generate token and refersh token (generate all tokens function from helper)

	token, refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.FirstName, *user.LastName, user.UserID)
	user.Token = &token
	user.RefreshToken = &refreshToken
	//if all ok, then you insert this new user into the user collection

	resultInsertionNumber, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer cancel()
	//return status OK and send the result back

	c.JSON(http.StatusOK, resultInsertionNumber)
}

func LoginUser(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User
	defer cancel()

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//find a user with that email and see if that user even exists
	var existingUser models.User
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)

	if err != nil {
		c.JSON(500, gin.H{"error": "User does not exist"})
		return
	}

	//verify the given password against the user password in database
	errHash := utils.CompareHashPassword(*user.Password, *existingUser.Password)

	if !errHash {
		c.JSON(400, gin.H{"error": "Invalid password"})
		return
	}

	//generate token for successfully authenticated user
	token, refreshToken, _ := helper.GenerateAllTokens(*existingUser.Email, *existingUser.FirstName, *existingUser.LastName, existingUser.UserID)

	//update tokens
	helper.UpdateAllTokens(token, refreshToken, existingUser.UserID)

	c.JSON(http.StatusOK, existingUser)

	//you can also set as a cookie if you want

}
