package app

import (
	"context"
	"grpc-todo/pkg/api/models"
	"grpc-todo/pkg/pb"
	"net/http"
)

func (s *Server) List(ctx context.Context, req *pb.Empty) (*pb.ListResponse, error) {
	var err error
	var todos []models.Todo
	var todosResponse []*pb.Todo

	if err = s.todoService.GetTodos(&todos); err != nil {
		return &pb.ListResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	for _, t := range todos {
		todosResponse = append(todosResponse, &pb.Todo{
			Content: t.Content,
			UserID:  uint64(t.UserID),
			Status:  t.Status,
		})
	}

	return &pb.ListResponse{
		Status: http.StatusOK,
		Error:  "",
		Todos:  todosResponse,
	}, nil
}

func (s *Server) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	var err error
	todo := new(models.Todo)

	todo.ID = uint(req.Id)

	if err = s.todoService.GetTodo(todo); err != nil {
		return &pb.GetTodoResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	return &pb.GetTodoResponse{
		Status: http.StatusOK,
		Error:  "",
		Todo: &pb.Todo{
			Content: todo.Content,
			Status:  todo.Status,
			UserID:  uint64(todo.UserID),
		},
	}, nil
}

func (s *Server) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	var err error

	todo := models.Todo{Content: req.Content, Status: req.Status, UserID: uint(req.UserID)}
	if err = s.todoService.CreateTodo(todo); err != nil {
		return &pb.CreateTodoResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	return &pb.CreateTodoResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}

func (s *Server) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.UpdateTodoResponse, error) {
	var err error

	todo := models.Todo{Content: req.Content, Status: req.Status}
	if err = s.todoService.UpdateTodo(todo); err != nil {
		return &pb.UpdateTodoResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	return &pb.UpdateTodoResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}

func (s *Server) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	var err error

	if err = s.todoService.DeleteTodo(uint(req.Id)); err != nil {
		return &pb.DeleteTodoResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	return &pb.DeleteTodoResponse{
		Status: http.StatusOK,
		Error:  "",
	}, nil
}
