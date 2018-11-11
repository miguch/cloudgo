package service

import (
	"fmt"
	view "github.com/miguch/cloudgo/cloudgo-view"
	"github.com/urfave/negroni"
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
	templateRouter := view.NewTemplateRouter(serv.webRoot)

	n.UseHandler(templateRouter)

	n.Run(fmt.Sprintf("%v:%v", serv.address, serv.port))
}