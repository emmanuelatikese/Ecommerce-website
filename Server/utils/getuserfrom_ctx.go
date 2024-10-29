package response

import (
	"api/models"
	"log"
	"net/http"
)

func GetUserFromContext(r *http.Request) (models.User){
	res := r.Context().Value("User").(models.User)
	log.Println(res)
	return res
}