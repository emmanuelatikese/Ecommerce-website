package response

import (
	"api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchAndDeleteInCart(list []models.Cart, productId string) ([]models.Cart){
	if productId != ""{
		pri_ProductId, err := primitive.ObjectIDFromHex(productId)
		if err != nil{
			return nil
		}
		for i := range list{
			if list[i].ProductId == pri_ProductId{
				list = append(list[:i], list[i+1:]...)
				return list
			}
		}
	}
	return []models.Cart{}
}