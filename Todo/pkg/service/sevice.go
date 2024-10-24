package service

import (
	"github.com/dusnila/todo-app.git"
	"github.com/dusnila/todo-app.git/pkg/repository"
)

type Authorizatiom interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
}

type TodoItem struct {
}

type Service struct {
	Authorizatiom
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorizatiom: NewAuthService(repos.Authorizatiom),
	}
}
