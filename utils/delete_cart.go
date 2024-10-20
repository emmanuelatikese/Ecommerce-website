package response

import (
	"api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchAndDeleteInCart(list []models.Cart, productId primitive.ObjectID) ([]models.Cart){
	if productId.Hex() != ""{
		for i := range list{
			if list[i].ProductId == productId{
				list = append(list[:i], list[i+1:]...)
				return list
			}
		}
	}
	return []models.Cart{}
}