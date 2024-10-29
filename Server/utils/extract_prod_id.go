package response

import (
	"api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProdId(cartItem []models.Cart)[]primitive.ObjectID{
	var prodList []primitive.ObjectID
	for i := range cartItem{
		prodList = append(prodList, cartItem[i].ProductId)
	}
	return prodList
}