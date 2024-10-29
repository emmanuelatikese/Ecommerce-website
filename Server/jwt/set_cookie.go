package jwt_util

import (
	"net/http"
	"time"
)

func SetCookie(accessToken string, refreshToken string, w http.ResponseWriter) {
	accessCookies := &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(time.Minute * 60 * 15),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	}

	refreshCookies := &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	}
	http.SetCookie(w, accessCookies)
	http.SetCookie(w, refreshCookies)
}

func SetAccessCookie(accessToken string, w http.ResponseWriter){
	accessCookies := &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(time.Minute * 60 * 15),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	}
	http.SetCookie(w, accessCookies)
}