package repository

import (
	"github.com/dusnila/todo-app.git"
	"github.com/jmoiron/sqlx"
)

type Authorizatiom interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
}

type TodoItem struct {
}

type Repository struct {
	Authorizatiom
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorizatiom: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
