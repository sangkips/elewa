package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"first_name" bson:"first_name, omitempty"`
	LastName     *string            `json:"last_name" bson:"last_name, omitempty"`
	Password     *string            `json:"password"`
	Email        *string            `json:"email" bson:"email,unique=true"`
	PhoneNumber  *string            `json:"phone_number" bson:"phone_number, omitempty, unique=true"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json:"refresh_token"`
	UserID       string             `json:"user_id"`
}
