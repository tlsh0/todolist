package repo

import (
	"database/sql"
	"fmt"

	"github.com/tlsh0/todolist/internal/models"
)

type Repo interface {
	CreateUser(*models.User) error
	UpdateUser(*models.User) error
	DeleteUser(*models.User) error
	GetUserById(int) (*models.User, error)
	CreateTask(*models.Task) error
	UpdateTask(*models.Task) error
	DeleteTask(*models.Task) error
	GetTaskById(int) (*models.Task, error)
}

type repo struct {
	db *sql.DB
}

func New(db *sql.DB) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) CreateUser() error {
	fmt.Println("working...")
	return nil
}

func (r *repo) UpdateUser() error {
	return nil
}

func (r *repo) DeleteUser() error {
	return nil
}

func (r *repo) GetUserById() (*models.User, error) {
	return &models.User{}, nil
}

func (r *repo) CreateTask() error {
	return nil
}

func (r *repo) UpdateTask() error {
	return nil
}

func (r *repo) DeleteTask() error {
	return nil
}

func (r *repo) GetAllTasks() ([]*models.Task, error) {
	return nil, nil
}

func (r *repo) GetTaskById() (*models.Task, error) {
	return &models.Task{}, nil
}
