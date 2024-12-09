package handlers

import (
	"net/http"

	"github.com/tlsh0/todolist/internal/middleware"
)

func Routes() *http.ServeMux {
	routes := http.NewServeMux()
	routes.HandleFunc("POST /login", makeHTTPHandlerFunc(loginHandler))
	routes.HandleFunc("POST /logout", middleware.WithJWT(makeHTTPHandlerFunc(logoutHandler)))
	routes.HandleFunc("POST /signup", makeHTTPHandlerFunc(singupHandler))
	routes.HandleFunc("/tasks", middleware.WithJWT(makeHTTPHandlerFunc(tasksHandler)))
	routes.HandleFunc("/task/{id}", middleware.WithJWT(makeHTTPHandlerFunc(singleTaskHandler)))
	return routes
}
