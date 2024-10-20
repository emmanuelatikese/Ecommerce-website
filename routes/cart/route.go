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
	getCart_handler := http.HandlerFunc(cart_controller.GetCart)

	//routes
	mux.Handle("/cart/addCart", protect_middleware(addCart_handler)).Methods("PATCH")
	mux.Handle("/cart/remove", protect_middleware(removalCart_handler)).Methods("DELETE")
	mux.Handle("/cart/updateQty/{product_id}", protect_middleware(updateQty_handler)).Methods("PATCH")
	mux.Handle("/cart/getCart", protect_middleware(getCart_handler)).Methods("GET")
}