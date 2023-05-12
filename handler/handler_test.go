package handler_test

import (
	"errors"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"mqbased-contracts/pact"
	"testing"
)

type Car struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
}

var tesla = Car{
	ID: "2", Title: "Tesla", Color: "Red",
}

func Test_generateContract(t *testing.T) {
	pact := pact.CreatePact()

	msg := pact.AddMessage()
	msg.
		Given("given a message to PBB's lambda").
		ExpectsToReceive("no error").
		WithContent(tesla).
		AsType(&Car{})
	err := pact.VerifyMessageConsumer(t, msg, handlerWrapper)
	if err != nil {
		t.Fatalf("Error on Verify: %v", err)
	}

	// specify PACT publisher
	publisher := dsl.Publisher{}
	err = publisher.Publish(types.PublishRequest{
		PactURLs:        []string{"./pacts/sample_consumer_handler-sample_provider_handler.json"},
		PactBroker:      "https://pen.pactflow.io/", //link to remote Contract broker
		BrokerToken:     "jEQnxw7xWgYRv-3-G7Cx-g",   //PactFlow token
		ConsumerVersion: "2.0.1",
		Tags:            []string{"latest"},
	})
	if err != nil {
		t.Fatal(err)
	}
}

var handlerWrapper = func(m dsl.Message) error {
	return userHandler(*m.Content.(*Car))
}

var userHandler = func(car Car) error {
	if car.ID != "0" {
		return nil
	} else {
		return errors.New("no ID found")
	}
}
