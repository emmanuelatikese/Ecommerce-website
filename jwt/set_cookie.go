package jwt_util

import (
	"net/http"
	"time"
)

func SetCookie(access_token string, refresh_token string, w http.ResponseWriter) {
	access_cookies := &http.Cookie{
		Name:     "access_token",
		Value:    access_token,
		Expires:  time.Now().Add(time.Minute * 60 * 15),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	}

	refresh_cookies := &http.Cookie{
		Name:     "refresh_token",
		Value:    refresh_token,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	}
	http.SetCookie(w, access_cookies)
	http.SetCookie(w, refresh_cookies)
}

func SetAccessCookie(access_token string, w http.ResponseWriter){
	access_cookies := &http.Cookie{
		Name:     "access_token",
		Value:    access_token,
		Expires:  time.Now().Add(time.Minute * 60 * 15),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	}
	http.SetCookie(w, access_cookies)
}