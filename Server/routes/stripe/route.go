package stripe_route

import (
	"api/middlewares"
	stripe_section "api/stripe"
	"net/http"

	"github.com/gorilla/mux"
)

func StripeRoute(mux *mux.Router){
	protectorMiddlewareHandler := middlewares.ProtectRoute
	mux.Handle("/stripe/checkout",
	protectorMiddlewareHandler(http.HandlerFunc(
		stripe_section.CheckoutSession))).Methods("POST")

	mux.Handle("/stripe/success",
	protectorMiddlewareHandler(http.HandlerFunc(
		stripe_section.CheckoutSuccess))).Methods("POST")
}