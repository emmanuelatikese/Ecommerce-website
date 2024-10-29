package response

import (
	"api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchAndChangeQty(list []models.Cart, productId primitive.ObjectID, quantity int64) ([]models.Cart){
	for i := range list{
		if list[i].ProductId == productId{
			if quantity > 0{
			list[i].Quantity = quantity
			return list
			} else{
				list = append(list[:i], list[i+1:]...)
				return list
			}
		}
	}
	return list
}