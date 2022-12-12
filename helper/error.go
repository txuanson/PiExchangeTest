package helper

import (
	"testing"

	"github.com/tebeka/atexit"
)


func TestErrorHandler(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Encountered an error: %v", err)
	}
}

// HandleError is a helper function to handle errors using the atexit package.
// Couldn't find a way to test this function
func HandleError(err error) {
	if err != nil {
		atexit.Fatalf("Encountered an error: %v", err)
	}
}
