package product_controllers

import (
	db_mongo "api/db/mongo"
	response "api/utils"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllProduct(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	product_collection := db_mongo.ProductCollection
	result_cursor, err := product_collection.Find(ctx, bson.D{{}})
	if err != nil {
		http.Error(w, "Internal server Error", 500)
		return
	}
	defer result_cursor.Close(ctx)
	var all_products []bson.M
	err = result_cursor.All(ctx, &all_products)
	if err != nil {
		http.Error(w,  err.Error(), 500)
		return
	}
	if len(all_products) == 0 {
		http.Error(w, "No product", 404)
		return
	}
	response.JsonResponse(all_products, w, 200)
}
