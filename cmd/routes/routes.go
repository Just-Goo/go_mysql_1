package routes

import (
	"net/http"

	"github.com/Just-Goo/Go-MySql-1/cmd/handlers"
)

// Router returns the default http server mux
func Router() *http.ServeMux {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/home", handlers.HomeHandler)
	http.HandleFunc("/insert", handlers.InsertHandler)
	http.HandleFunc("/error", handlers.ErrorHandler)

	return http.DefaultServeMux
}
