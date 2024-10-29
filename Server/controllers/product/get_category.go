package product_controllers

import (
	db_mongo "api/db/mongo"
	response "api/utils"
	"net/http"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCategory(w http.ResponseWriter, r *http.Request){
	category := mux.Vars(r)["category"]
	if category == ""{
		http.Error(w, "No category was passed", 404)
		return
	}
	ctx, productCollection := r.Context(), db_mongo.ProductCollection
	var allCategory []bson.M
	cur, err := productCollection.Find(ctx, bson.M{"category": category})
	isErr := response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	err = cur.All(ctx, &allCategory);
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	if len(allCategory) == 0 {
		http.Error(w, "No product fall under this category", 404)
		return
	}
	response.JsonResponse(allCategory, w, 200)
}