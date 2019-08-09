// +build integration

package consumer

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pact-foundation/pact-go/dsl"
)

func TestCreateToDo(t *testing.T) {
	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Consumer: "ToDoConsumer",
		Provider: "ToDoService",
		Host:     "localhost",
	}
	defer pact.Teardown()

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
}
