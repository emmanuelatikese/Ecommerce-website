package stripe_section

import "net/http"

func CheckoutSuccess(w http.ResponseWriter, r *http.Request){
	sessionId := r.URL.Query().Get("session_id")
	if sessionId == ""{
		http.Error(w, "session_id not provided", 404)
		return
	}
}