package todo

import (
	"context"
	"github.com/go-pg/pg"
	"github.com/golang/protobuf/ptypes"
	"github.com/satori/go.uuid"
	"github.com/xuanit/testing/todo/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Service is the service dealing with storing
// and retrieving pb items from the database.
type Service struct {
	DB *pg.DB
}

// CreateTodo creates a pb given a description
func (s Service) CreateTodo(ctx context.Context, req *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	id, _ := uuid.NewV4()
	req.Item.Id = id.String()
	err := s.DB.Insert(req.Item)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not insert item into the database: %s", err)
	}
	return &pb.CreateTodoResponse{Id: req.Item.Id}, nil
}

// CreateTodos create pb items from a list of pb descriptions
func (s Service) CreateTodos(ctx context.Context, req *pb.CreateTodosRequest) (*pb.CreateTodosResponse, error) {
	var ids []string
	for _, item := range req.Items {
		id, _ := uuid.NewV4()
		item.Id = id.String()
		ids = append(ids, item.Id)
	}
	err := s.DB.Insert(&req.Items)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not insert items into the database: %s", err)
	}
	return &pb.CreateTodosResponse{Ids: ids}, nil
}

// GetTodo retrieves a pb item from its ID
func (s Service) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	var item pb.Todo
	err := s.DB.Model(&item).Where("id = ?", req.Id).First()
	if err != nil {
		return nil, grpc.Errorf(codes.NotFound, "Could not retrieve item from the database: %s", err)
	}
	return &pb.GetTodoResponse{Item: &item}, nil
}

// ListTodo retrieves a pb item from its ID
func (s Service) ListTodo(ctx context.Context, req *pb.ListTodoRequest) (*pb.ListTodoResponse, error) {
	var items []*pb.Todo
	query := s.DB.Model(&items).Order("created_at ASC")
	if req.Limit > 0 {
		query.Limit(int(req.Limit))
	}
	if req.NotCompleted {
		query.Where("completed = false")
	}
	err := query.Select()
	if err != nil {
		return nil, grpc.Errorf(codes.NotFound, "Could not list items from the database: %s", err)
	}
	return &pb.ListTodoResponse{Items: items}, nil
}

// DeleteTodo deletes a pb given an ID
func (s Service) DeleteTodo(ctx context.Context, req *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	err := s.DB.Delete(&pb.Todo{Id: req.Id})
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not delete item from the database: %s", err)
	}
	return &pb.DeleteTodoResponse{}, nil
}

// UpdateTodo updates a pb item
func (s Service) UpdateTodo(ctx context.Context, req *pb.UpdateTodoRequest) (*pb.UpdateTodoResponse, error) {
	req.Item.UpdatedAt = ptypes.TimestampNow()
	res, err := s.DB.Model(req.Item).Column("title", "description", "completed", "updated_at").Update()
	if res.RowsAffected() == 0 {
		return nil, grpc.Errorf(codes.NotFound, "Could not update item: not found")
	}
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not update item from the database: %s", err)
	}
	return &pb.UpdateTodoResponse{}, nil
}

// UpdateTodos updates pb items given their respective title and description.
func (s Service) UpdateTodos(ctx context.Context, req *pb.UpdateTodosRequest) (*pb.UpdateTodosResponse, error) {
	time := ptypes.TimestampNow()
	for _, item := range req.Items {
		item.UpdatedAt = time
	}
	res, err := s.DB.Model(&req.Items).Column("title", "description", "completed", "updated_at").Update()
	if res.RowsAffected() == 0 {
		return nil, grpc.Errorf(codes.NotFound, "Could not update items: not found")
	}
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "Could not update items from the database: %s", err)
	}
	return &pb.UpdateTodosResponse{}, nil
}
