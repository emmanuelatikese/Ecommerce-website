package middlewares

import (
	response "api/utils"
	"log"
	"net/http"
)

func AdminProtect(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		log.Println("here admin")
		user, ok := response.GetUserFromContext(r)
		if !ok{
			http.Error(w, "Internal Server error", 500)
			return
		}
		if user == nil{
			http.Error(w, "Empty", 500)
			return
		}
		if user["role"] != "admin"{
			http.Error(w, "Admin only!", http.StatusUnauthorized)
			return
		}
		log.Println("there there")
		next.ServeHTTP(w,r)
	})
}