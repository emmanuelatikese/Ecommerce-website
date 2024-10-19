package cart_controller

import (
	db_mongo "api/db/mongo"
	"api/models"
	response "api/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RemoveProductCart(w http.ResponseWriter, r *http.Request){
	var product_id models.ProductId
	user, ok:= response.GetUserFromContext(r)
	if !ok {
		http.Error(w, "Unauthorized", 500)
		return
	}
	ctx, user_collection := r.Context(), db_mongo.UserCollection
	userStr, ok := user["_id"].(string)
	if !ok{
		http.Error(w, "Unauthorized", 500)
		return
	}
	pri_userId, err := primitive.ObjectIDFromHex(userStr)
	if err != nil{
		http.Error(w, err.Error(), 500)
		return
	}
	pri_ProductId, err := primitive.ObjectIDFromHex(product_id.Id)
	if err != nil{
		http.Error(w, err.Error(), 500)
		return
	}
	user_cartItem, ok := user["cart_item"].([]models.Cart)
	if !ok{
		http.Error(w, "Internal server error", 500)
		return
	}
	var filterUser bson.M
	new_cart := response.SearchAndDeleteInCart(user_cartItem, pri_ProductId)

	err = user_collection.FindOneAndUpdate(ctx, bson.M{"_id": pri_userId},
		bson.M{"$set":bson.M{"cart_item": new_cart}}, 
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&filterUser)
	if err != nil{
			http.Error(w, err.Error(), 500)
			return
		}
	response.JsonResponse(filterUser["cart_item"], w, 200)
}