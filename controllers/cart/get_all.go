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
	ctx, productCollection := r.Context(), db_mongo.ProductCollection
	prodList := response.GetProdId(user.CartItem)
	cursor, err := productCollection.Find(ctx, bson.M{"_id": bson.M{"$in": prodList}})
	isErr := response.ErrorHandler(err, w,500)
	if isErr{
		return
	}
	var FilterProducts  []models.Product
	err = cursor.All(ctx, &FilterProducts)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	userProducts := response.GetPrdQty(FilterProducts, user.CartItem)
	response.JsonResponse(userProducts, w, 200)
}