package cart_controller

import (
	db_mongo "api/db/mongo"
	"api/models"
	response "api/utils"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCart(w http.ResponseWriter, r *http.Request){
	user := response.GetUserFromContext(r)
	ctx, product_collection := r.Context(), db_mongo.ProductCollection
	prodList := response.GetProdId(user.CartItem)
	cursor, err := product_collection.Find(ctx, bson.M{"_id": bson.M{"$in": prodList}})
	if err != nil{
		http.Error(w, err.Error(), 500)
		return
	}
	var FilterProducts  []models.Product
	if err := cursor.All(ctx, &FilterProducts); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	user_products := response.GetPrdQty(FilterProducts, user.CartItem)
	response.JsonResponse(user_products, w, 200)
}