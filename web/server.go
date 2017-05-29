package web

import (
	"fmt"
	"net/http"
)

func NewServer() {
	var router *http.ServeMux

	router = http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome home!")
	})

	http.ListenAndServe(":8000", router)
}
