package repository

import (
	"github.com/go-pg/pg"
	"github.com/xuanit/testing/todo/pb"
)

type ToDoImpl struct {
	DB *pg.DB
}

type ToDo interface {
	List(limit int32, notCompleted bool) ([]*pb.Todo, error)
	Insert(items *pb.Todo)  error
	Get(id string) (*pb.Todo, error)
	Delete(id string) error
}

func (r ToDoImpl) List(limit int32, notCompleted bool) ([]*pb.Todo, error) {
	var items []*pb.Todo
	query := r.DB.Model(&items).Order("created_at ASC")
	if limit > 0 {
		query.Limit(int(limit))
	}
	if notCompleted {
		query.Where("completed = false")
	}
	err := query.Select()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r ToDoImpl) Insert(items *pb.Todo)  error {
	err := r.DB.Insert(items)
	return err
}

func (r ToDoImpl) Get(id string) (*pb.Todo, error) {
		var item pb.Todo
		err := r.DB.Model(&item).Where("id = ?", id).First()
		if err != nil {
			return nil, err
		}
		return &item, nil
}

func (r ToDoImpl) Delete(id string) error {
	err := r.DB.Delete(&pb.Todo{Id: id})
	return err
}
