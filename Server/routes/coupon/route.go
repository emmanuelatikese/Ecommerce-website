package coupon_route

import (
	coupon_controllers "api/controllers/coupon"
	"api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func CouponRoute(mux *mux.Router){
	protectMiddleware := middlewares.ProtectRoute
	mux.Handle("/coupon/get", protectMiddleware(http.HandlerFunc(coupon_controllers.GetCoupons))).Methods("GET")
	mux.Handle("/coupon/validate", protectMiddleware(http.HandlerFunc(coupon_controllers.ValidateCoupon))).Methods("GET")
}