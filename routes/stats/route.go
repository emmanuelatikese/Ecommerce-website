package stats_route

import (
	stats_controller "api/controllers/stats"
	"api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func StatsRoute(mux *mux.Router){
		adminMiddleware, protectMiddleware := middlewares.AdminProtect, middlewares.ProtectRoute
		getAnalyticsHandler := http.HandlerFunc(stats_controller.GetAnalytics)
		mux.Handle("/stats/getAnalytics", protectMiddleware(
			adminMiddleware(getAnalyticsHandler))).Methods("GET")
		
}