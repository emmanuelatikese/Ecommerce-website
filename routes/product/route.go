package product_route

import (
	product_controllers "api/controllers/product"
	"api/middlewares"
	"net/http"
	"github.com/gorilla/mux"
)

func ProductRoute(mux *mux.Router){
	admin_middleware, protect_middleware := middlewares.AdminProtect, middlewares.ProtectRoute
	get_all_handler := http.HandlerFunc(product_controllers.GetAllProduct)
	mux.Handle("/product/getAll", protect_middleware(admin_middleware(get_all_handler))).Methods("GET")
	mux.HandleFunc("/product/getFeatured", product_controllers.GetFeaturedProduct).Methods("GET")
}