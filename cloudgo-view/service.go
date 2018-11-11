package cloudgo_view

import (
	"github.com/gorilla/mux"
	"github.com/kataras/go-template"
	"github.com/kataras/go-template/pug"
	"github.com/miguch/cloudgo/cloudgo-entity"
	"net/http"
)


func NewTemplateRouter(root string) (*mux.Router) {
	//Pug template engine
	template.AddEngine(pug.New(pug.Config{})).Directory(root + "/cloudgo-view/template", ".pug")
	err := template.Load()
	if err != nil {
		panic(err)
	}

	renderSignup := func(w http.ResponseWriter, r *http.Request) {
		//Execute template engine, pass in a map as template parameters
		err := template.ExecuteWriter(w, "signup.pug", map[string]interface{}{"message": ""})
		if err != nil {
			w.Write([]byte(err.Error()))
		}
	}

	renderSignin := func(w http.ResponseWriter, r *http.Request) {
		vars := r.URL.Query()
		message := ""
		username, ok := vars["username"]
		if ok {
			//Render the user info page if user exists
			if user := cloudgo_entity.GetUser(username[0]); user != nil {
				err := template.ExecuteWriter(w, "userInfoPage.pug", map[string]interface{}{
					"username": user.Username,
					"studentID": user.StudentID,
					"phoneNumber": user.Phone,
					"email": user.Email,
					"message": "",
				})
				if err != nil {
					w.Write([]byte(err.Error()))
				}
				return
			} else {
				message = "Queried user does not exist!"
			}
		}
		//Render the login page if no user specified or user not exists
		err := template.ExecuteWriter(w, "signin.pug", map[string]interface{}{"message": message})
		if err != nil {
			w.Write([]byte(err.Error()))
		}
	}

	//Use Gorilla mux as router
	router := mux.NewRouter()

	//Render login or info page when GET /
	router.HandleFunc("/", renderSignin).Methods("GET")

	//render the registration page if GET /regist
	router.HandleFunc("/regist", renderSignup).Methods("GET")

	return router
}

