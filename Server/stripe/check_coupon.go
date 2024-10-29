package stripe_section

import (
	"api/models"
	response "api/utils"
	"net/http"
	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/coupon"
)

func CheckCouponForDiscount(filterCoupon models.Coupons, w http.ResponseWriter)string{
	NewCoupon := &stripe.CouponParams{
		Duration: stripe.String(string(stripe.CouponDurationOnce)),
		PercentOff: stripe.Float64(filterCoupon.Discount),
	}
	couponStripe, err := coupon.New(NewCoupon)
	isErr := response.ErrorHandler(err, w, 500)
	if isErr{
		return ""
	}
	return couponStripe.ID
}