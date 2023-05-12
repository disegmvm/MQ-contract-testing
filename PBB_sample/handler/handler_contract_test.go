package handler

import (
	"github.com/pact-foundation/pact-go/dsl"
	"mqbased-contracts/pact"
	"path/filepath"
	"testing"
)

var tesla = Car{
	ID: "2", Title: "Tesla", Color: "Red",
}

func Test_verifyContract(t *testing.T) {
	pactClient := pact.CreatePact()

	// Map test descriptions to message producer (handlers)
	functionMappings := dsl.MessageHandlers{
		// Description (ExpectedToReceive)
		"no error": func(m dsl.Message) (interface{}, error) {
			if tesla.ID != "0" {
				return tesla, nil
			} else {
				return map[string]string{
					"message": "not found",
				}, nil
			}
		},
	}

	stateMappings := dsl.StateHandlers{
		// Provider state (Given)
		"given a message to PBB's lambda": func(s dsl.State) error {
			// Content and Type
			return nil
		},
	}

	// Verify the Provider with local Pact Files
	_, err := pactClient.VerifyMessageProvider(t, dsl.VerifyMessageRequest{
		BrokerURL:                  "https://pen.pactflow.io", //link to remote Contract broker
		BrokerToken:                "jEQnxw7xWgYRv-3-G7Cx-g",  //PactFlow token
		PublishVerificationResults: true,
		//ProviderTags:               []string{"latest v2"},
		ProviderVersion: "1.100.0",
		MessageHandlers: functionMappings,
		StateHandlers:   stateMappings,
		PactLogDir:      filepath.ToSlash("./logs"),
		PactLogLevel:    "DEBUG",
	})
	if err != nil {
		return
	}
}
