package product_route

import (
	product_controllers "api/controllers/product"
	"api/middlewares"
	"net/http"
	"github.com/gorilla/mux"
)

func ProductRoute(mux *mux.Router){
	//Handlers
	adminMiddleware, protectMiddleware := middlewares.AdminProtect, middlewares.ProtectRoute
	getAllHandler := http.HandlerFunc(product_controllers.GetAllProduct)
	createProductHandler := http.HandlerFunc(product_controllers.CreateProduct)
	deleteProductHandler := http.HandlerFunc(product_controllers.DeleteProduct)
	updateFeaturedHandler := http.HandlerFunc(product_controllers.UpdateFeatured)
	getRecommendationHandler := http.HandlerFunc(product_controllers.GetRecommendation)
	getCategoryHandler := http.HandlerFunc(product_controllers.GetCategory)
	// deleteAllHandler := http.HandlerFunc(product_controllers.DeleteAll)
	//Routes
	mux.Handle("/product/getAll", protectMiddleware(adminMiddleware(getAllHandler))).Methods("GET")
	mux.HandleFunc("/product/getFeatured", product_controllers.GetFeaturedProduct).Methods("GET")
	mux.Handle("/product/create", protectMiddleware(adminMiddleware(createProductHandler))).Methods("POST")
	mux.Handle("/product/delete/{id}", protectMiddleware(adminMiddleware(deleteProductHandler))).Methods("DELETE")
	mux.Handle("/product/recommendation", protectMiddleware(adminMiddleware(getRecommendationHandler))).Methods("GET")
	mux.Handle("/product/category", protectMiddleware(adminMiddleware(getCategoryHandler))).Methods("GET")
	mux.Handle("/product/updateFeatured/{id}", protectMiddleware(adminMiddleware(updateFeaturedHandler))).Methods("PATCH")
	// mux.Handle("/product/deleteAll", protect_middleware(admin_middleware(deleteAllHandler))).Methods("DELETE")
}