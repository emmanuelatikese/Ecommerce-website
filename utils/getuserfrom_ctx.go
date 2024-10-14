package response

import (
	"log"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserFromContext(r *http.Request) (bson.M, bool){
	res := r.Context().Value("User").(bson.M)
	if res == nil {
		log.Fatal("Empty")
		return nil, false
	}
	return res, true
}