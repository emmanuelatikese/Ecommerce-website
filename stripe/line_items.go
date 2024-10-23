package stripe_section

import (
	"api/models"

	"github.com/stripe/stripe-go/v80"
)

func LineItemsAndTotalAmt(prod []models.ProductQty) ([]*stripe.CheckoutSessionLineItemParams, float64){
	var totalAmount float64
	var allList  []*stripe.CheckoutSessionLineItemParams

	for i := range prod {
		amount := prod[i].Products.Price
		totalAmount = amount * 100 // changing to cents
		item := &stripe.CheckoutSessionLineItemParams{
         PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
                    Currency: stripe.String("usd"),
                    ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
                        Name: stripe.String(prod[i].Products.Name),
						Description: stripe.String(prod[i].Products.Description),
						Images: []*string{stripe.String(prod[i].Products.Image)},
                    },
                    UnitAmount: stripe.Int64(int64(amount)),
                },
      }
	  allList = append(allList, item)
	}

	return allList, totalAmount
}