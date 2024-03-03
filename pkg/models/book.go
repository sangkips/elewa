package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	Price      *float64           `json:"price" bson:"price, omitempty"`
	Book_id    string             `json:"book_id" bson:"book_id"`
	CategoryId *string            `json:"category_id" bson:"category_id, omitempty"`
	Author     *Author
}

type Author struct {
	FirstName string `json:"first_name,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"last_name,omitempty" bson:"lastname,omitempty"`
}
