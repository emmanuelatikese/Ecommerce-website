package cart_controller

import (
	db_mongo "api/db/mongo"
	"api/models"
	response "api/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func UpdateQuantity(w http.ResponseWriter, r *http.Request){
	product_id := mux.Vars(r)["product_id"]
	log.Print(product_id, "here")
	user:= response.GetUserFromContext(r)
	var ProdQty models.ProductIdQty
	err := json.NewDecoder(r.Body).Decode(&ProdQty)
	isErr := response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	ctx, userCollection := r.Context(), db_mongo.UserCollection
	priProductId, err := primitive.ObjectIDFromHex(product_id)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	var filterUser bson.M
	new_cart := response.SearchAndChangeQty(user.CartItem, priProductId, ProdQty.Quantity)
	err = userCollection.FindOneAndUpdate(ctx, bson.M{"_id": user.Id},
		bson.M{"$set":bson.M{"cart_item": new_cart}},
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&filterUser)
	isErr = response.ErrorHandler(err, w, 404)
	if isErr{
		return
	}
	response.JsonResponse(filterUser["cart_item"], w, 200)
}