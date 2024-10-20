package carts_route

import (
	cart_controller "api/controllers/cart"
	"api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func CartRoute(mux *mux.Router){
	protect_middleware := middlewares.ProtectRoute
	addCart_handler := http.HandlerFunc(cart_controller.AddToCart)
	removalCart_handler := http.HandlerFunc(cart_controller.RemoveProductCart)
	updateQty_handler := http.HandlerFunc(cart_controller.UpdateQuantity)
	//GetAll_handler
	
	//routes
	mux.Handle("/cart/create", protect_middleware(addCart_handler)).Methods("PATCH")
	mux.Handle("/cart/remove/{id}", protect_middleware(removalCart_handler)).Methods("PATCH")
	mux.Handle("/cart/updateQuality/{id}", protect_middleware(updateQty_handler)).Methods("PATCH")

}