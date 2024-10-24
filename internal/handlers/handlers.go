package handlers

import (
	"net/http"

	"github.com/tlsh0/todolist/internal/models"
)

type Data struct {
	User  *models.User
	Tasks []*models.Task
}

func Home(w http.ResponseWriter, r *http.Request) {
	// TODO: list the tasks or redirect to login page
}

func HomePost(w http.ResponseWriter, r *http.Request) {
	// TODO: update the tasks
}

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: handle both GET method to DISPLAY the page and to LOGIN
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	// TODO: handle both POST method to DISPLAY the page and to LOGIN
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: logout the user using method POST
}
