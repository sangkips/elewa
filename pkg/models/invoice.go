package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	InvoiceId      string             `json:"invoice_id"`
	OrderId        string             `json:"order_id"`
	PaymentMethod  *string            `json:"payment_method" bson:"payment_method" vlaidate:"eq=CARD|eq=CASH|eq=MPESA"`
	PaymentStatus  *string            `json:"payment_status" bson:"payment_status" validate:"required, eq=PENDING|eq=PAID"`
	PaymentDueDate time.Time          `json:"payment_due_date"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at"`
}
