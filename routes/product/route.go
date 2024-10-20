package product_route

import (
	product_controllers "api/controllers/product"
	"api/middlewares"
	"net/http"
	"github.com/gorilla/mux"
)

func ProductRoute(mux *mux.Router){
	//Handlers
	admin_middleware, protect_middleware := middlewares.AdminProtect, middlewares.ProtectRoute
	getAll_handler := http.HandlerFunc(product_controllers.GetAllProduct)
	createProduct_handler := http.HandlerFunc(product_controllers.CreateProduct)
	deleteProduct_handler := http.HandlerFunc(product_controllers.DeleteProduct)
	updateFeatured_handler := http.HandlerFunc(product_controllers.UpdateFeatured)
	getRecommendation_handler := http.HandlerFunc(product_controllers.GetRecommendation)
	getCategory_handler := http.HandlerFunc(product_controllers.GetCategory)
	deleteAll_handler := http.HandlerFunc(product_controllers.DeleteAll)
	//Routes
	mux.Handle("/product/getAll", protect_middleware(admin_middleware(getAll_handler))).Methods("GET")
	mux.HandleFunc("/product/getFeatured", product_controllers.GetFeaturedProduct).Methods("GET")
	mux.Handle("/product/create", protect_middleware(admin_middleware(createProduct_handler))).Methods("POST")
	mux.Handle("/product/delete/{id}", protect_middleware(admin_middleware(deleteProduct_handler))).Methods("DELETE")
	mux.Handle("/product/recommendation", protect_middleware(admin_middleware(getRecommendation_handler))).Methods("GET")
	mux.Handle("/product/category", protect_middleware(admin_middleware(getCategory_handler))).Methods("GET")
	mux.Handle("/product/updateFeatured/{id}", protect_middleware(admin_middleware(updateFeatured_handler))).Methods("PATCH")
	mux.Handle("/product/deleteAll", protect_middleware(admin_middleware(deleteAll_handler))).Methods("DELETE")

}