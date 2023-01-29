package app

import (
	"grpc-todo/pkg/api/services"
	"grpc-todo/pkg/pb"
)

type Server struct {
	todoService services.TodoService
	pb.UnimplementedTodoServiceServer
}

func NewServer(todoService services.TodoService) *Server {
	return &Server{todoService: todoService}
}
