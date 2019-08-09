// +build integration

package consumer

import (
	"log"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/stretchr/testify/assert"
)

func TestToDoProxy(t *testing.T) {
	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Consumer: "ToDoConsumer",
		Provider: "ToDoService",
		Host:     "localhost",
	}
	defer pact.Teardown()

	t.Run("TestCreateToDo", func(t *testing.T) {
		createToDoRes := struct {
			id string `json:"id" pact:"example=id1"`
		}{}
		// Set up our expected interactions.
		pact.
			AddInteraction().
			//Given("UserA is existing").
			UponReceiving("A request to create todo").
			WithRequest(dsl.Request{
				Method:  "POST",
				Path:    dsl.String("/v1/todo"),
				Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
				Body: map[string]interface{}{
					"title":       "1-1 with manager",
					"description": "discuss about OKRs",
					"completed":   true,
				},
			}).
			WillRespondWith(dsl.Response{
				Status:  200,
				Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
				Body:    dsl.Match(createToDoRes),
			})

		// Pass in test case. This is the component that makes the external HTTP call
		var test = func() (err error) {
			proxy := ToDoProxy{Host: "localhost", Port: pact.Server.Port}
			id, err := proxy.CreateToDo(ToDo{Id: "1", Title: "1-1 with manager", Description: "discuss about OKRs", Completed: true})
			if err != nil {
				return err
			}
			assert.Equal(t, "id1", id)
			return nil
		}

		// Run the test, verify it did what we expected and capture the contract
		if err := pact.Verify(test); err != nil {
			log.Fatalf("Error on Verify: %v", err)
		}
	})

	t.Run("TestGetToDoList", func(t *testing.T) {
		pact.AddInteraction().Given("There are todo A and todo B").UponReceiving("A request to list todo").
			WithRequest(dsl.Request{
				Method: "GET",
				Path:   dsl.String("/v1/todo"),
				Query:  dsl.MapMatcher{"limit": dsl.String("3"), "not_completed": dsl.String("true")},
			}).WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body: dsl.Like(map[string]interface{}{
				"items": dsl.Like([]interface{}{
					dsl.Like(map[string]interface{}{"id": dsl.Like("id1"), "title": dsl.Like("ToDo A")}),
					dsl.Like(map[string]interface{}{"id": dsl.Like("id2"), "title": dsl.Like("ToDo B")}),
				}),
			}),
		})

		test := func() (err error) {
			proxy := ToDoProxy{Host: "localhost", Port: pact.Server.Port}
			_, err = proxy.ListToDo(3, true)

			if err != nil {
				return err
			}
			return nil
		}

		if err := pact.Verify(test); err != nil {
			log.Fatalf("Error on Verify: %v", err)
		}
	})
}
