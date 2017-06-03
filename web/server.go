package web

import (
	"fmt"
	"github.com/PauloASilva/hello-world/config"
	"github.com/PauloASilva/hello-world/controllers"
	"github.com/PauloASilva/m1"
	"net/http"
)

// Creates and starts an HTTP server listening at conf.Port
func NewServer(conf *config.Config) {
	var router *http.ServeMux
	var listenAddress string

	router = http.NewServeMux()

	// Register application controllers
	controllers.HomeRegister(router)
	controllers.AboutRegister(router)

	// Load modules
	m1.Register(router)

	// start server
	listenAddress = fmt.Sprintf(":%d", conf.Port)
	fmt.Printf("Server listening at %s\n", listenAddress)
	http.ListenAndServe(listenAddress, router)
}
