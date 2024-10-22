package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID	`bson:"_id,omitempty"`
	Username string 			`json:"username"`
	Password []byte 			`json:"password"`
	Email    string 			`json:"email"`
	CartItem []Cart 			`json:"cartitem"` // quantity- 1 as default
	Role string 				`json:"role"` //enum either customer or admin but customer is default
}

type Cart struct {
	ProductId primitive.ObjectID	`json:"product_id"`
	Quantity float64				`json:"quantity"`
}