package services

import (
	"errors"
	"grpc-todo/pkg/api/models"
)

type TodoRepository interface {
	Get(todos *[]models.Todo) error
	GetOne(todo *models.Todo) error
	Create(todo models.Todo) error
	Update(todo models.Todo) error
	Delete(id uint) error
}

type TodoService interface {
	GetTodos(todos *[]models.Todo) error
	GetTodo(todo *models.Todo) error
	CreateTodo(todo models.Todo) error
	UpdateTodo(todo models.Todo) error
	DeleteTodo(id uint) error
}

type todoService struct {
	storage TodoRepository
}

func NewTodoService(repository TodoRepository) TodoService {
	return &todoService{storage: repository}
}

func (t *todoService) GetTodos(todos *[]models.Todo) error {
	var err error
	if err = t.storage.Get(todos); err != nil {
		return err
	}

	return nil
}

func (t *todoService) GetTodo(todo *models.Todo) error {
	var err error

	if todo.ID == 0 {
		return errors.New("server[todo]-id required")
	}

	if err = t.storage.GetOne(todo); err != nil {
		return err
	}

	return nil
}

func (t *todoService) CreateTodo(todo models.Todo) error {
	var err error

	if todo.Content == "" {
		return errors.New("server[todo]-content required")
	}

	if todo.UserID == 0 {
		return errors.New("server[todo]-userID required")
	}

	if err = t.storage.Create(todo); err != nil {
		return err
	}

	return nil
}

func (t *todoService) UpdateTodo(todo models.Todo) error {
	var err error
	if err = t.storage.Update(todo); err != nil {
		return err
	}

	return nil
}

func (t *todoService) DeleteTodo(id uint) error {
	var err error

	if id == 0 {
		return errors.New("server[todo]-id required")
	}

	if err = t.storage.Delete(id); err != nil {
		return err
	}

	return nil
}
