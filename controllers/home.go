package controllers

import (
	"fmt"
	"net/http"
)

// Binds controller's actions to routes
func HomeRegister(router *http.ServeMux) {
	router.HandleFunc("/", home)
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
