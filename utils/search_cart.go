package response

import (
	"api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchAndUpdateInCart(list []models.Cart, productId primitive.ObjectID) ([]models.Cart){
	if productId.Hex() != ""{
		for i := range list{
			if list[i].ProductId == productId{
				list[i].Quantity++
				return list
			}
		}
	}
	list = append(list, models.Cart{ProductId: productId, Quantity: 1})
	return list
}