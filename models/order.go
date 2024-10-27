package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	Id      		primitive.ObjectID		`bson:"_id,omitempty"`
	User    		User					`json:"user"`
	Products 		[]ProductQty			`json:"product"`
	TotalAmount 	float64					`json:"totalamount"`
	StripeSessionId string					`json:"sessionid"`
}