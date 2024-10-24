package router

import (
	"net/http"

	"github.com/tlsh0/todolist/internal/handlers"
)

func NewRouter() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /", handlers.Home)
	r.HandleFunc("POST /", handlers.Home)
	r.HandleFunc("GET /login", handlers.Login)
	r.HandleFunc("POST /login", handlers.LoginPost)
	r.HandleFunc("POST /logout", handlers.Logout)

	return r
}
