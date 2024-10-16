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
	ctx, product_collection := r.Context(), db_mongo.ProductCollection
	var allCategory []bson.M
	cur, err := product_collection.Find(ctx, bson.M{"category": category})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err = cur.All(ctx, &allCategory); err != nil{
		http.Error(w,err.Error(), 500)
		return
	}

	if len(allCategory) == 0 {
		http.Error(w, "No product fall under this category", 404)
		return
	}
	response.JsonResponse(allCategory, w, 200)
}