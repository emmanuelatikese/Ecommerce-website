package stripe_section

import (
	"api/models"

	"github.com/stripe/stripe-go/v80"
)

func LineItemsAndTotalAmt(prod []models.ProductQty) ([]*stripe.CheckoutSessionLineItemParams, float64){
	var totalAmount float64
	var allList  []*stripe.CheckoutSessionLineItemParams

	for i := range prod {
		amount := prod[i].Product.Price* 100  // changing to cents
		totalAmount = amount * float64(prod[i].Quantity)
		item := &stripe.CheckoutSessionLineItemParams{
         PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
                    Currency: stripe.String("usd"),
                    ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
                        Name: stripe.String(prod[i].Product.Name),
						Description: stripe.String(prod[i].Product.Description),
						// Images: []*string{stripe.String(prod[i].Product.Image)},
                    },
                    UnitAmount: stripe.Int64(int64(amount)),
                },
		Quantity: stripe.Int64(prod[i].Quantity),
      }
	  allList = append(allList, item)
	}

	return allList, totalAmount
}