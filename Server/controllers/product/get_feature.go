package product_controllers

import (
	db_mongo "api/db/mongo"
	redis_db "api/db/redis"
	response "api/utils"
	"net/http"
	"time"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFeaturedProduct(w http.ResponseWriter, r *http.Request){
	var isFeatured []bson.M
	ctx := r.Context()
	productCollection := db_mongo.ProductCollection
	cursor, err := productCollection.Find(ctx, bson.M{"isfeatured": true})
	isErr := response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &isFeatured);
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	if len(isFeatured) == 0 {
		http.Error(w, "No featured products available", 500)
		return
	}
	redis_db.RedisCli.Set(ctx, "featured_product", isFeatured, time.Hour * 24 * 7)
	response.JsonResponse(isFeatured, w, 200)
}