package product_controllers

import (
	db_mongo "api/db/mongo"
	response "api/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

)

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	ctx, productCollection := r.Context(), db_mongo.ProductCollection
	alltask ,err := productCollection.DeleteMany(ctx,  bson.D{{}})
	response.ErrorHandler(err, w, 500)
	if alltask.DeletedCount == 0 {
		response.JsonResponse("No tasks available", w, 404)
		return
	}
	response.JsonResponse("products deleted successfully", w, 200)
}