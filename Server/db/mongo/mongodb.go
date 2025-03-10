package db_mongo

import (
	"context"
	"os"

	"github.com/subosito/gotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	UserCollection    *mongo.Collection
	ProductCollection *mongo.Collection
	CouponCollection  *mongo.Collection
	OrderCollection   *mongo.Collection
)


func ConnectDB() error{
	gotenv.Load()
	url := os.Getenv("MONGO_URL")
	ctx := context.TODO()
	mongo_url := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, mongo_url)
	if err != nil {
		return err
	}
	db := client.Database("E-commence_db")
	UserCollection = db.Collection("user-collection")
	ProductCollection = db.Collection("product-collection")
	CouponCollection = db.Collection("coupon-collection")
	OrderCollection = db.Collection("order-collection")
	return nil
}