package cloudgo_entity

import (
	"net/http"
	"strconv"
)

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	//Get username and password from POST parameter
	name := r.Form.Get("username")
	pass := r.Form.Get("password")
	res := CheckSignin(name, pass)
	if res == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(strconv.Itoa(res)))
	}
}


