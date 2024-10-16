package product_controllers

import (
	db_mongo "api/db/mongo"
	redis_db "api/db/redis"
	response "api/utils"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetFeaturedProduct(w http.ResponseWriter, r *http.Request){
	var isFeatured []bson.M
	ctx := r.Context()
	product_collection := db_mongo.ProductCollection
	cursor, err := product_collection.Find(ctx, bson.M{"is_featured": true})
	if err != nil{
		http.Error(w, err.Error(), 500)
		return
	}

	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &isFeatured); err != nil{
		log.Println("here")
		http.Error(w, err.Error(), 500)
		return
	}

	if len(isFeatured) == 0 {
		http.Error(w, "No product or no featured product", 500)
		return
	}
	redis_db.RedisCli.Set(ctx, "featured_product", isFeatured, time.Hour * 24 * 7)
	response.JsonResponse(isFeatured, w, 200)
}