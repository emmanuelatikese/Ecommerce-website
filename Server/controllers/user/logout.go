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
	isErr := response.ErrorHandler(err, w, http.StatusUnauthorized)
	if isErr{
		return
	}
	val := jwt.Value
	userId, err := jwt_util.Verify(w,r,val, jwt_util.RefreshTokenPublic)
	if err != nil || userId == ""{
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	redis_db.RedisCli.Del(ctx, userId)
	accessCookies := &http.Cookie{
		Name: "access_token",
		Value: "",
		Path:     "/",
		Expires: time.Unix(0,0),
        HttpOnly: true,
        Secure:   true,
        MaxAge:   -1,
	}
	refreshCookies := &http.Cookie{
		Name: "refresh_token",
		Value: "",
		Path:     "/",
		Expires: time.Unix(0,0),
        HttpOnly: true,
        Secure:   true,
        MaxAge:   -1,
	}

	http.SetCookie(w, accessCookies)
	http.SetCookie(w, refreshCookies)
	response.JsonResponse("Logout successfully..", w, 200)
}