package user_route

import (
	user_controllers "api/controllers/user"
	"api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRoutes (mux *mux.Router) {
	mux.HandleFunc("/auth/SignUp", user_controllers.SignUp).Methods("POST")
	mux.HandleFunc("/auth/Login", user_controllers.Login).Methods("POST")
	mux.HandleFunc("/auth/Logout", user_controllers.Logout).Methods("POST")
	mux.HandleFunc("/auth/Refresh_token", user_controllers.RefreshAccess).Methods("POST")
	mux.Handle("/user/profile", (middlewares.ProtectRoute(
		http.HandlerFunc(user_controllers.GetProfile)))).Methods("POST")
}