package response

import (
	"api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchAndChangeQty(list []models.Cart, productId primitive.ObjectID, quantity int) ([]models.Cart){
	if productId.Hex() != ""{
		for i := range list{
		if list[i].ProductId == productId{
			list[i].Quantity = quantity
			return list
		}
		}
	}
	return []models.Cart{}
}