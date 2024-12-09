package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type customHandler func(w http.ResponseWriter, r *http.Request) error

func loginHandler(w http.ResponseWriter, r *http.Request) error {
	var req = new(userCredentials)
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return writeJSON(w, http.StatusBadRequest, errorResponse{Error: "bad request"})
	}

	//TODO: login the user and asign a jwt token

	return nil
}

func logoutHandler(w http.ResponseWriter, r *http.Request) error {
	// token := r.Header.Get("x-jwt-token")
	//TODO: validate the token and set expiration time to 0.
	return writeJSON(w, http.StatusOK, successResponse{Success: "logged out successfully"})
}

func singupHandler(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func tasksHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet: // list the tasks
	case http.MethodPatch: // rearrange the tasks
	case http.MethodPost: // add a new task
	}
	return writeJSON(w, http.StatusMethodNotAllowed, fmt.Sprintf("method not allowed: %s", r.Method))
}

func singleTaskHandler(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodPatch: // change the status of the task
	case http.MethodDelete: // delete the task completely
	}
	return writeJSON(w, http.StatusMethodNotAllowed, fmt.Sprintf("method not allowed: %s", r.Method))
}

func makeHTTPHandlerFunc(f customHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error()})
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
