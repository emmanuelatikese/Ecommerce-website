package response

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(content interface{}, w http.ResponseWriter, code int){
	str_json, err := json.Marshal(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type","application/json")
	w.WriteHeader(code)
	w.Write(str_json)
}