package main

import (
	"github.com/miguch/cloudgo/service"
	flag "github.com/spf13/pflag"
)

const (
	PORT = 8080
	ADDRESS = "127.0.0.1"
	WEBROOT = "."
)

func main() {
	port := flag.Uint16P("port", "p", PORT, "The port number which the program will listen on")
	address := flag.StringP("address", "a", ADDRESS, "The address which the program will listen on")
	webRoot := flag.StringP("webroot", "r", WEBROOT, "The root of the application")
	flag.Parse()

	server := service.NewServer(*address, *port, *webRoot)
	server.Run()
}
