package response

import (
	"api/models"
	"log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchAndUpdateInCart(list []models.Cart, productId primitive.ObjectID) ([]models.Cart){
	log.Println(len(list))
	if len(list) > 0{
		for i := range list{
			if list[i].ProductId == productId{
				log.Println("here")
				list[i].Quantity++
				return list
			}
		}
	}
	list = append(list, models.Cart{ProductId: productId, Quantity: 1})
	return list
}