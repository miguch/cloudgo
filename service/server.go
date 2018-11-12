package service

import (
	"fmt"
	"github.com/miguch/cloudgo/cloudgo-entity"
	view "github.com/miguch/cloudgo/cloudgo-view"
	"github.com/urfave/negroni"
	"net/http"
)

type Server struct {
	address string
	port uint16
	webRoot string
}

func NewServer(address string, port uint16, webRoot string) (*Server) {
	return &Server{
		address,
		port,
		webRoot,
	}
}

func (serv *Server) Run() {
	n := negroni.Classic()

	n.UseFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if r.Method == "POST" {
			r.ParseForm()
		}
		next.ServeHTTP(w, r)
	})

	router := view.NewTemplateRouter(serv.webRoot)

	//process regist info if POST /regist
	router.HandleFunc("/regist", cloudgo_entity.SignupHandler).Methods("POST")
	//process login info
	router.HandleFunc("/signin", cloudgo_entity.SigninHandler).Methods("POST")

	//static file server
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(serv.webRoot + "/cloudgo-view/assets")))

	//Use router in negroni
	n.UseHandler(router)

	n.Run(fmt.Sprintf("%v:%v", serv.address, serv.port))
}