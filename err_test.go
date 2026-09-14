package common

import (
	"fmt"
	"testing"
)

func ExampleZeroError() {
	realErr := ZeroError(exampleError)
	noErr := ZeroError(exampleNoError)
	goErr := ZeroError(exampleGoError)
	diag := ZeroError(exampleDiag)

	fmt.Printf("Real error: %t\n", realErr) 	   // Result: 66 (not allowed on non-leaf)
	fmt.Printf("Go error:   %t\n", goErr)   	   // any non-nil standard Go error
	fmt.Printf("No error:   %t\n", noErr)   	   // Result: 0 (success)
	fmt.Printf("Diag only:  %t\n", diag)    	   // Result: 0 (success, but only contains a diagnostic message)
	fmt.Printf("Idiomatic:  %t\n", exampleDiag == nil) // inappropriate idiomatic eval.
	// Output:
	// Real error: false
	// Go error:   false
	// No error:   true
	// Diag only:  true
	// Idiomatic:  false
}

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

var (
	exampleError error = LDAPResultNotAllowedOnNonLeaf.New( "Operation not allowed on non-leaf")
	exampleNoError error = LDAPResultSuccess.New()
	exampleGoError error = fmt.Errorf("A standard Go error")
	exampleDiag error
)


func init() {
	diag := LDAPResultSuccess.New().(Error)
	diag.SetDiag("Diagnostic message")
	exampleDiag = diag
}
