package response

import "net/http"

func ErrorHandler(err error, w http.ResponseWriter, code int) bool{
	if err != nil{
		http.Error(w, err.Error(), code)
		return true
	}
	return false
}