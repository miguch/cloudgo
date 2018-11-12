package cloudgo_entity

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	registMode := r.Form.Get("register")
	if registMode == "true" {
		returnCode := 1
		name := r.Form.Get("username")
		if !CheckUsernameDuplicate(name) {
			returnCode *= 3
		}
		email := r.Form.Get("mail")
		if !CheckEmailDuplicate(email) {
			returnCode *= 11
		}
		id := r.Form.Get("number")
		if !CheckIDDuplicate(id) {
			returnCode *= 4
		}
		phone := r.Form.Get("phone")
		if !CheckPhoneDuplicate(phone) {
			returnCode *= 7
		}
		pass := r.Form.Get("password")
		if returnCode != 1 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(strconv.Itoa(returnCode)))
		} else {
			err := addUser(User{
				name,
				id,
				email,
				phone,
				pass,
			})
			if err != nil {
				fmt.Fprintf(os.Stderr,"Error Adding user: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		}
	} else if registMode == "false" {
		result := true
		name := r.Form.Get("username")
		if name != "" {
			result = CheckUsernameDuplicate(name)
		}
		email := r.Form.Get("mail")
		if email != "" {
			result = CheckEmailDuplicate(email)
		}
		id := r.Form.Get("number")
		if id != "" {
			result = CheckIDDuplicate(id)
		}
		phone := r.Form.Get("phone")
		if phone != "" {
			result = CheckPhoneDuplicate(phone)
		}
		if !result {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

}