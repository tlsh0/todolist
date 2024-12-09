package models

import "time"

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
