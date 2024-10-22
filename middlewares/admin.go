package middlewares

import (
	response "api/utils"
	"net/http"
)

func AdminProtect(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		user := response.GetUserFromContext(r)
		if user.Role != "admin"{
			http.Error(w, "Admin only!", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w,r)
	})
}