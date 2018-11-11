package cloudgo_view

import (
	"github.com/gorilla/mux"
	"github.com/kataras/go-template"
	"github.com/kataras/go-template/pug"
	"github.com/miguch/cloudgo/cloudgo-entity"
	"net/http"
)


func NewTemplateRouter(root string) (*mux.Router) {
	template.AddEngine(pug.New(pug.Config{})).Directory(root + "/cloudgo-view/template", ".pug")
	err := template.Load()
	if err != nil {
		panic(err)
	}

	renderSignup := func(w http.ResponseWriter, r *http.Request) {
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
		err := template.ExecuteWriter(w, "signin.pug", map[string]interface{}{"message": message})
		if err != nil {
			w.Write([]byte(err.Error()))
		}
	}

	router := mux.NewRouter()
	router.HandleFunc("/", renderSignin)
	router.HandleFunc("/regist", renderSignup).Methods("GET")
	router.HandleFunc("/regist", cloudgo_entity.SignupHandler).Methods("POST")
	router.HandleFunc("/signin", cloudgo_entity.SigninHandler).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir(root + "/cloudgo-view/assets")))

	return router
}

