package product_controllers

import (
	db_mongo "api/db/mongo"
	response "api/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

)

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	ctx, product_collection := r.Context(), db_mongo.ProductCollection
	alltask ,err := product_collection.DeleteMany(ctx,  bson.D{{}})
	if err !=nil{
		http.Error(w, err.Error(), 500)
		return
	}
	if alltask.DeletedCount == 0 {
		response.JsonResponse("No tasks available", w, 404)
		return
	}
	response.JsonResponse("products deleted successfully", w, 200)
}