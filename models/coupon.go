package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Coupons struct {
	Id			primitive.ObjectID	`bson:"_id,omitempty"`
	Code      	string  			`json:"code"`
	Discount 	float64 			`json:"discount"`
	Expiration	time.Time 			`json:"expiration"`
	IsActive	bool 				`json:"isactive"`
	UserId 		primitive.ObjectID 	`json:"userid"`
	CreatedAt	time.Time
	UpdatedAt 	time.Time
}

type ProdCoupon struct{
	Products []ProductQty `json:"products"`
	Coupon  string		`json:"coupon"`
}
