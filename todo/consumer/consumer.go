package consumer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang/protobuf/ptypes/timestamp"
)

type ToDo struct {
	Id          string               `json:"id,omitempty" `
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

type ToDoList struct {
	Items []ToDo `json:"items"`
}

func (p *ToDoProxy) getApi() string {
	return fmt.Sprintf("http://%s:%d/v1", p.Host, p.Port)
}

func (p *ToDoProxy) CreateToDo(todo ToDo) (string, error) {
	todo.Id = ""
	toDoBytes, _ := json.Marshal(todo)

	u := fmt.Sprintf("%s/todo", p.getApi())
	req, _ := http.NewRequest("POST", u, bytes.NewBuffer(toDoBytes))
	req.Header.Set("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)

	createToDoRes := struct {
		ID string `json:"id"`
	}{}
	json.NewDecoder(res.Body).Decode(&createToDoRes)

	res.Body.Close()
	return createToDoRes.ID, nil
}

func (p *ToDoProxy) ListToDo(limit int32, notCompleted bool) (*ToDoList, error) {
	url := fmt.Sprintf("%s/todo?limit=%d&not_completed=%t", p.getApi(), limit, notCompleted)
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	res, _ := http.DefaultClient.Do(req)
	toDos := ToDoList{}
	json.NewDecoder(res.Body).Decode(toDos)

	res.Body.Close()
	return &toDos, nil
}
