package response

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUserFromContext(r *http.Request) (bson.M, bool){
	return r.Context().Value("User").(bson.M), true
}