package user_controllers

import (
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request){
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
}