package cart_controller

import (
	db_mongo "api/db/mongo"
	"api/models"
	response "api/utils"
	"encoding/json"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RemoveProductCart(w http.ResponseWriter, r *http.Request){
	var product_id models.ProductIdQty
	err := json.NewDecoder(r.Body).Decode(&product_id)
	response.ErrorHandler(err, w, http.StatusBadRequest)
	user:= response.GetUserFromContext(r)
	if len(user.CartItem) == 0{
		response.JsonResponse("Cart is already empty", w, http.StatusNotAcceptable)
		return
	}
	ctx, user_collection := r.Context(), db_mongo.UserCollection
	var filterUser bson.M
	new_cart := response.SearchAndDeleteInCart(user.CartItem, product_id.Id)
	err = user_collection.FindOneAndUpdate(ctx, bson.M{"_id": user.Id},
		bson.M{"$set":bson.M{"cartitem": new_cart}}, 
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&filterUser)
	response.ErrorHandler(err, w, 404)
	response.JsonResponse(filterUser["cartitem"], w, 200)
}