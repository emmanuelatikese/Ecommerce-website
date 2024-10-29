package carts_route

import (
	cart_controller "api/controllers/cart"
	"api/middlewares"
	"net/http"
	"github.com/gorilla/mux"
)

func CartRoute(mux *mux.Router){
	protectMiddleware := middlewares.ProtectRoute
	addCartHandler := http.HandlerFunc(cart_controller.AddToCart)
	removalCartHandler := http.HandlerFunc(cart_controller.RemoveProductCart)
	updateQtyHandler := http.HandlerFunc(cart_controller.UpdateQuantity)
	getCartHandler := http.HandlerFunc(cart_controller.GetCart)

	//routes
	mux.Handle("/cart/addCart", protectMiddleware(addCartHandler)).Methods("PATCH")
	mux.Handle("/cart/remove", protectMiddleware(removalCartHandler)).Methods("DELETE")
	mux.Handle("/cart/updateQty/{product_id}", protectMiddleware(updateQtyHandler)).Methods("PATCH")
	mux.Handle("/cart/getCart", protectMiddleware(getCartHandler)).Methods("GET")
}