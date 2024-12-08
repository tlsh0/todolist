package models

import "time"

type NewTaskRequest struct {
	Name    string `json:"name"`
	DueTime string `json:"due_time"`
}

type DeleteTaskRequest struct {
	ID int64 `json:"id"`
}

type UpdateTaskRequest struct {
	ID   int64 `json:"id"`
	Done bool  `json:"status"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID                int64
	Username          string
	Email             string
	EncryptedPassword string
}

type Task struct {
	ID      int64
	UserID  int64
	Name    string
	DueTime time.Time
	Done    bool
	Order   int64
}
