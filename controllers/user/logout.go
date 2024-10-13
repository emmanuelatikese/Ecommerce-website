package user_controllers

import (
	redis_db "api/db/redis"
	jwt_util "api/jwt"
	response "api/utils"
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	jwt, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	val := jwt.Value
	userId, err := jwt_util.Verify(w,r,val, jwt_util.Refresh_token_public)
	if err != nil || userId == ""{
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	redis_db.RedisCli.Del(ctx, userId)
	access_cookies := &http.Cookie{
		Name: "access_token",
		Value: "",
		Path:     "/",
		Expires: time.Unix(0,0),
        HttpOnly: true,
        Secure:   true,
        MaxAge:   -1,
	}
	refresh_cookies := &http.Cookie{
		Name: "refresh_token",
		Value: "",
		Path:     "/",
		Expires: time.Unix(0,0),
        HttpOnly: true,
        Secure:   true,
        MaxAge:   -1,
	}

	http.SetCookie(w, access_cookies)
	http.SetCookie(w, refresh_cookies)
	response.JsonResponse("Logout successfully..", w, 200)
}