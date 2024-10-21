package response

import "api/models"

func GetPrdQty(list []models.Product, cartItems []models.Cart) []models.ProductQty{
	var prod_quantity []models.ProductQty
	for i := range cartItems{
		for x := range list{
			if cartItems[i].ProductId == list[x].Id{
				prod_quantity = append(prod_quantity, models.ProductQty{Products: list[x],
					Quantity: cartItems[i].Quantity})
			}
		}
	}
	return prod_quantity
}