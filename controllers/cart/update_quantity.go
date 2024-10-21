package cart_controller

import (
	db_mongo "api/db/mongo"
	"api/models"
	response "api/utils"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func UpdateQuantity(w http.ResponseWriter, r *http.Request){
	product_id := mux.Vars(r)["product_id"]
	user:= response.GetUserFromContext(r)

	var ProdQty models.ProductIdQty
	err := json.NewDecoder(r.Body).Decode(&ProdQty)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	ctx, user_collection := r.Context(), db_mongo.UserCollection

	pri_ProductId, err := primitive.ObjectIDFromHex(product_id)
	if err != nil{
		http.Error(w, err.Error(), 500)
		return
	}
	var filterUser bson.M
	new_cart := response.SearchAndChangeQty(user.CartItem, pri_ProductId, ProdQty.Quantity)
	err = user_collection.FindOneAndUpdate(ctx, bson.M{"_id": user.Id},
		bson.M{"$set":bson.M{"cart_item": new_cart}},
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&filterUser)
	if err != nil{
			http.Error(w, err.Error(), 404)
			return
		}
	response.JsonResponse(filterUser["cart_item"], w, 200)
}