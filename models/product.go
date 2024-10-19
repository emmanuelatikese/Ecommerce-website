package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Name string `json:"name"`
	Price int `json:"price"`
	Description string `json:"description"`
	Image string `json:"image"`
	Category string `json:"category"`
	IsFeatured bool `json:"is_featured"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductId struct{
	Id string `json:"product_id"`
}