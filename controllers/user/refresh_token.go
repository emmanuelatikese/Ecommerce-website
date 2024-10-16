package user_controllers

import (
	redis_db "api/db/redis"
	jwt_util "api/jwt"
	response "api/utils"
	"net/http"
)

func RefreshAccess(w http.ResponseWriter, r *http.Request) {
	jwt_cookies, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ctx := r.Context()
	jwt_token := jwt_cookies.Value

	userId, err := jwt_util.Verify(w, r, jwt_token, jwt_util.Refresh_token_public)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	refresh_redis_token, err:= redis_db.RedisCli.Get(ctx, userId).Result()
	if err != nil{
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

	if refresh_redis_token != jwt_token{
		http.Error(w, "Refresh token Invalid", http.StatusNotAcceptable)
		return
	}
	access_token, err := jwt_util.GenerateAccessToken(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jwt_util.SetAccessCookie(access_token, w)
	response.JsonResponse("Token refreshed successfully", w, http.StatusCreated)
}