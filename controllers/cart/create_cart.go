package cart_controller

import (
	db_mongo "api/db/mongo"
	"api/models"
	response "api/utils"
	"encoding/json"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo/options"
)


func AddToCart(w http.ResponseWriter, r *http.Request){
	var productId models.ProductIdQty
	ctx, user_collection := r.Context(), db_mongo.UserCollection
	err := json.NewDecoder(r.Body).Decode(&productId)
	response.ErrorHandler(err, w, http.StatusBadRequest)
	user := response.GetUserFromContext(r)
	pri_ProductId, err := primitive.ObjectIDFromHex(productId.Id)
	response.ErrorHandler(err, w, 500)
	updateCart := response.SearchAndUpdateInCart(user.CartItem, pri_ProductId)
	var filterUser bson.M
	_, err = user_collection.UpdateOne(ctx, bson.M{"_id": user.Id},
		bson.M{"$set":bson.M{"cartitem": updateCart}},)
	response.ErrorHandler(err, w, 500)
	err = user_collection.FindOne(ctx, bson.M{"_id": user.Id}).Decode(&filterUser)
	response.ErrorHandler(err, w, 500)
	response.JsonResponse(filterUser["cartitem"], w, 200)
}