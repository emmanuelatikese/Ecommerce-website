package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`
	Price int `json:"price"`
	Image string `json:"image"`
	Category string `json:"category"`
	IsFeatured bool `json:"is_featured"`
}