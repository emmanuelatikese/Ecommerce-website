package cart_controller

import (
	db_mongo "api/db/mongo"
	"api/models"
	response "api/utils"
	"log"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCart(w http.ResponseWriter, r *http.Request){
	log.Println("here")
	user := response.GetUserFromContext(r)
	ctx, product_collection := r.Context(), db_mongo.ProductCollection
	prodList := response.GetProdId(user.CartItem)
	cursor, err := product_collection.Find(ctx, bson.M{"_id": bson.M{"$in": prodList}})
	response.ErrorHandler(err, w,500)
	var FilterProducts  []models.Product
	err = cursor.All(ctx, &FilterProducts)
	response.ErrorHandler(err, w, 500)
	user_products := response.GetPrdQty(FilterProducts, user.CartItem)
	response.JsonResponse(user_products, w, 200)
}