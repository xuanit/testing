package service

import (
	"context"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/xuanit/testing/todo/pb"
	"github.com/xuanit/testing/todo/server/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// ToDo is the service dealing with storing
// and retrieving pb items from the database.
type ToDo struct {
	//DB *pg.DB
	ToDoRepo repository.ToDo
}

// CreateTodo creates a pb given a description
func (s ToDo) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	id, _ := uuid.NewV4()
	req.Item.Id = id.String()
	err := s.ToDoRepo.Insert(req.Item)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not insert item into the database: %s", err)
	}
	return &pb.CreateTodoResponse{Id: req.Item.Id}, nil
}

// GetTodo retrieves a pb item from its ID
func (s ToDo) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	item, err := s.ToDoRepo.Get(req.Id)
	if err != nil {
		return nil, grpc.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", err)
	}
	return &pb.GetTodoResponse{Item: item}, nil
}

// ListTodo retrieves a pb item from its ID
func (s ToDo) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoResponse, error) {
	var items []*pb.Todo
	items, err := s.ToDoRepo.List(req.Limit, req.NotCompleted)
	fmt.Printf("------------- List TODO  %v------\n", items)
	if err != nil {
		return nil, grpc.Errorf(codes.NotFound, "Could not list items from the database: %s", err)
	}
	return &pb.ListTodoResponse{Items: items}, nil
}

// DeleteTodo deletes a pb given an ID
func (s ToDo) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	err := s.ToDoRepo.Delete(req.Id)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not delete item from the database: %s", err)
	}
	return &pb.DeleteTodoResponse{}, nil
}
