package stripe_section

import (
	response "api/utils"
	"encoding/json"
	"net/http"
)

func CheckoutSuccess(w http.ResponseWriter, r *http.Request){
	type newSession struct {
		Id string `json:"sessionId"`
	}
	var sessionId newSession
	err := json.NewDecoder(r.Body).Decode(&sessionId)
	response.ErrorHandler(err, w, http.StatusBadRequest)
	sessId := sessionId.Id
	if sessId == ""{
		http.Error(w, "No session Id provided", 404)
		return
	}
}