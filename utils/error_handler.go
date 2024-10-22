package response

import "net/http"

func ErrorHandler(err error, w http.ResponseWriter, code int){
	if err != nil{
		http.Error(w, err.Error(), code)
	}
}