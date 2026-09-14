package common

import (
	"fmt"
	"testing"
)

func ExampleError_idiomatic() {
	// pretend exampleError was just handed to you
	// after attempting to delete a subtree without
	// the appropriate control OID.
	if exampleError != nil {
		fmt.Println(exampleError)
	}
	// Output: Operation not allowed on non-leaf
}

func ExampleError_assertion() {
	// pretend exampleError was just handed to you,
	// but you want to see if it is a standard go
	// error or common.Error instance.
	switch err := exampleError.(type) {
	case Error:
		fmt.Printf("Code: %d\n", err.Code)
		fmt.Printf("Text: %s\n", err.Error())
	default:
		// not a common.Error
	}
	// Output:
	// Code: 66
	// Text: Operation not allowed on non-leaf
}

func ExampleErrorIs() {
	// pretend exampleError was just handed to you,
	// but you wish to see if it is not only an
	// instance of Error, but also whether it
	// contains a specific result code (66).
	fmt.Println(ErrorIs(exampleError, 66))
	// Output: true
}

func TestErrorIs_codecov(t *testing.T) {
	_ = ErrorIs(nil, 0)
	newRes(0)
}

var exampleError error = newRes(66, "Operation not allowed on non-leaf")
