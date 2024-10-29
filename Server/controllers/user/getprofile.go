package user_controllers

import (
	response "api/utils"
	"net/http"
)

func GetProfile(w http.ResponseWriter, r *http.Request){
	user := response.GetUserFromContext(r)
	response.JsonResponse(user, w, 200)
}