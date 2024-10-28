package product_controllers

import (
	db_mongo "api/db/mongo"
	response "api/utils"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllProduct(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	productCollection := db_mongo.ProductCollection
	result_cursor, err := productCollection.Find(ctx, bson.D{{}})
	isErr := response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	defer result_cursor.Close(ctx)
	var allProducts []bson.M
	err = result_cursor.All(ctx, &allProducts)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	if len(allProducts) == 0 {
		http.Error(w, "No products available", 404)
		return
	}
	response.JsonResponse(allProducts, w, 200)
}
