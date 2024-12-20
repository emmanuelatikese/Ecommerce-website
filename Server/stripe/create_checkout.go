package stripe_section

import (
	db_mongo "api/db/mongo"
	"api/models"
	response "api/utils"
	"encoding/json"
	"net/http"
	"os"
	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/checkout/session"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/subosito/gotenv"
)

func CheckoutSession(w http.ResponseWriter, r *http.Request){
	gotenv.Load()
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	clientUrl := os.Getenv("CLIENT_URL")
	var req models.ProdCoupon

	err := json.NewDecoder(r.Body).Decode(&req)
	isErr := response.ErrorHandler(err, w, http.StatusBadRequest)
	if isErr{
		return
	}
	if len(req.Products) == 0 {
		http.Error(w, "No products", 404)
		return
	}
	strProd, err := json.Marshal(req.Products)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	ctx, couponCollection := r.Context(), db_mongo.CouponCollection
	user := response.GetUserFromContext(r)
	var coupon string

	var stripeCoupon stripe.CheckoutSessionDiscountParams
	lineItems, totalAmount := LineItemsAndTotalAmt(req.Products)

	if req.Coupon != ""{
		priId, err := primitive.ObjectIDFromHex(req.Coupon)
		isErr = response.ErrorHandler(err, w, 500)
		if isErr{
			return
		}
		var filterCoupon models.Coupons
		err = couponCollection.FindOne(ctx, bson.M{"_id": priId, "isactive": true, "userid":user.Id}).Decode(&filterCoupon)
		if err != nil {
			if err == mongo.ErrNoDocuments{
				http.Error(w, "coupon don't exist", 404)
				return
			}
			http.Error(w, err.Error(), 500)
			return
		}
		coupon = filterCoupon.Code
		totalAmount *= (100 - filterCoupon.Discount)/100
		newCoupon := CheckCouponForDiscount(filterCoupon, w)
		stripeCoupon = stripe.CheckoutSessionDiscountParams{Coupon: stripe.String(newCoupon)}
	}
	if coupon == ""{
		stripeCoupon = stripe.CheckoutSessionDiscountParams{}
	}

	successUrl := clientUrl + "/success/{checkoutSessionUrl}"
	cancelUrl := clientUrl + "/cancel"

	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
            "card",
        }),
		Mode: stripe.String("payment"),
		LineItems: lineItems,
		SuccessURL: stripe.String(successUrl),
        CancelURL:  stripe.String(cancelUrl),
		Discounts: []*stripe.CheckoutSessionDiscountParams{
			&stripeCoupon,
		},
		Metadata: map[string]string{
			"userId": *stripe.String(user.Id.Hex()),
			"couponCode":*stripe.String(req.Coupon),
			"products": *stripe.String(string(strProd)),
		},
	}

	newSession, err := session.New(params)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	if totalAmount >= 20000{
		CreateCoupon(user.Id, ctx, couponCollection,w)
	}
	response.JsonResponse(map[string]interface{}{
		"sessionId": newSession.ID,
		"totalAmount": totalAmount / 100,
	}, w, 201)

}