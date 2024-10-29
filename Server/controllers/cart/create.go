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
	ctx, userCollection := r.Context(), db_mongo.UserCollection
	err := json.NewDecoder(r.Body).Decode(&productId)

	isErr := response.ErrorHandler(err, w, http.StatusBadRequest)
	if isErr {
		return
	}
	user := response.GetUserFromContext(r)
	priProductId, err := primitive.ObjectIDFromHex(productId.Id)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{ 
		return
	}

	updateCart := response.SearchAndUpdateInCart(user.CartItem, priProductId)
	var filterUser bson.M
	_, err = userCollection.UpdateOne(ctx, bson.M{"_id": user.Id},
		bson.M{"$set":bson.M{"cartitem": updateCart}},)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr {
		return
	}
	err = userCollection.FindOne(ctx, bson.M{"_id": user.Id}).Decode(&filterUser)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	response.JsonResponse(filterUser["cartitem"], w, 200)
}