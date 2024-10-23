package models

import (
	"fmt"
	"time"
)

type User struct {
	ID       int64
	Username string
	Name     string
	Password string
	Email    string
	Tasks    []*Task
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %s>", u.ID, u.Username, u.Email)
}

type Task struct {
	TaskId    int64
	UserId    int64
	Name      string
	Done      bool
	Order     int
	CreatedAt time.Time
	ExpiresAt time.Time
}

func (t Task) String() string {
	return fmt.Sprintf("Task<%d %d %s %t %d \"%d:%d %d-%d-%d\" \"%d:%d %d-%d-%d\"\n", t.TaskId, t.UserId, t.Name, t.Done, t.Order, t.CreatedAt.Hour(), t.CreatedAt.Minute(), t.CreatedAt.Day(), t.CreatedAt.Month(), t.CreatedAt.Year(), t.ExpiresAt.Hour(), t.ExpiresAt.Minute(), t.ExpiresAt.Day(), t.ExpiresAt.Month(), t.ExpiresAt.Year())
}
