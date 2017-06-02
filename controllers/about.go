package controllers

import (
	"fmt"
	"net/http"
)

// Binds controller's actions to routes
func AboutRegister(router *http.ServeMux) {
	router.HandleFunc("/about", main)
}

func main(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is a Hello World Go WebApp!")
}
