package web

import (
	"fmt"
	"github.com/PauloASilva/hello-world/config"
	"github.com/PauloASilva/hello-world/controllers"
	"net/http"
	"strings"
)

type handler func(w http.ResponseWriter, req *http.Request)

type Server struct {
	config *config.Config
	router *http.ServeMux
}

func (s *Server) setConfig(conf *config.Config) {
	s.config = conf
}

// Initializes `server` router and registers controllers
func (s *Server) createRouter() {
	s.router = http.NewServeMux()

	controllers.HomeRegister(s.router)
	controllers.AboutRegister(s.router)
}

// Registers module's routes (`mRoutes`) under a base route `baseRoute`
func (s *Server) RegisterModule(baseRoute string,
	mRoutes map[string]func(w http.ResponseWriter, req *http.Request)) {
	for r := range mRoutes {
		route := fmt.Sprintf("%s/%s", baseRoute, strings.TrimPrefix(r, "/"))
		s.router.HandleFunc(route, mRoutes[r])
	}
}

// starts server
func (s *Server) Start() {
	listenAddress := fmt.Sprintf(":%d", s.config.Port)
	fmt.Printf("Server listening at %s\n", listenAddress)
	http.ListenAndServe(listenAddress, s.router)
}

// Creates and starts an HTTP server listening at conf.Port
func CreateServer(conf *config.Config) Server {
	var s Server

	s.setConfig(conf)
	s.createRouter()

	return s
}
