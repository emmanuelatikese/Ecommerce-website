package user_controllers

import (
	redis_db "api/db/redis"
	jwt_util "api/jwt"
	response "api/utils"
	"net/http"
)

func RefreshAccess(w http.ResponseWriter, r *http.Request) {
	jwtCookies, err := r.Cookie("refresh_token")
	isErr := response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	ctx := r.Context()
	jwtToken := jwtCookies.Value

	userId, err := jwt_util.Verify(w, r, jwtToken, jwt_util.RefreshTokenPublic)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	refresh_redis_token, err:= redis_db.RedisCli.Get(ctx, userId).Result()
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	if refresh_redis_token != jwtToken{
		http.Error(w, "Refresh token Invalid", http.StatusNotAcceptable)
		return
	}
	accessToken, err := jwt_util.GenerateAccessToken(userId)
	isErr = response.ErrorHandler(err, w, 500)
	if isErr{
		return
	}
	jwt_util.SetAccessCookie(accessToken, w)
	response.JsonResponse("Token refreshed successfully", w, http.StatusCreated)
}