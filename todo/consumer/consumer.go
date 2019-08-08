package consumer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang/protobuf/ptypes/timestamp"
)

type ToDo struct {
	Id          string               `json:"id,omitempty"`
	Title       string               `json:"title,omitempty"`
	Description string               `json:"description,omitempty"`
	Completed   bool                 `json:"completed,omitempty"`
	CreatedAt   *timestamp.Timestamp `json:"created_at,omitempty"`
	UpdatedAt   *timestamp.Timestamp `json:"updated_at,omitempty"`
}

type ToDoProxy struct {
	Host string
	Port int
}

func (p *ToDoProxy) getApi() string {
	return fmt.Sprintf("http://%s:%d/v1", p.Host, p.Port)
}

func (p *ToDoProxy) CreateToDo(todo ToDo) error {
	todo.Id = ""
	toDoBytes, err := json.Marshal(todo)

	u := fmt.Sprintf("%s/todo", p.getApi())
	req, _ := http.NewRequest("POST", u, bytes.NewBuffer(toDoBytes))
	req.Header.Set("Content-Type", "application/json")
	_, err = http.DefaultClient.Do(req)
	return err
}
