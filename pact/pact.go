package pact

import (
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"os"
)

var dir, _ = os.Getwd()
var logDir = fmt.Sprintf("%s/log", dir)

// Set up the Pact client
func CreatePact() dsl.Pact {
	return dsl.Pact{
		Consumer: "Sample consumer handler",
		Provider: "Sample provider handler",
		LogDir:   logDir,
		LogLevel: "INFO",
	}
}
