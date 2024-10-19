package cart_controller

import (
	db_mongo "api/db/mongo"
	"api/models"
	response "api/utils"
	"encoding/json"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func AddToCart(w http.ResponseWriter, r *http.Request){
	var productId string
	ctx, user_collection := r.Context(), db_mongo.UserCollection
	err := json.NewDecoder(r.Body).Decode(&productId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, ok := response.GetUserFromContext(r)
	if !ok{
		http.Error(w, "UnAuthorized", 500)
		return
	}
	userId := user["_id"]

	pri_ProductId, err := primitive.ObjectIDFromHex(productId)
	if err != nil{
		http.Error(w, err.Error(), 500)
		return
	}
	userCart, ok := user["cart_item"].([]models.Cart)
	if !ok{
		http.Error(w,"cart User can't convert", 500)
		return
	}
	updateCart := response.SearchAndUpdateInCart(userCart, pri_ProductId)
	var filterUser bson.M
	err = user_collection.FindOneAndUpdate(ctx, bson.M{"_id": userId},
		bson.M{"cart_item": updateCart}, 
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&filterUser)
	if err != nil{
			http.Error(w, err.Error(), 500)
			return
		}
	response.JsonResponse(filterUser["cart_item"], w, 200)
}