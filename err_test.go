package common

import (
	"fmt"
	"testing"
)

/*
This example demonstrates use of the [ZeroError] helper in multiple
scenarios.

A standard Go error will always return false if it is not nil according
to the standard idiomatic expression "err == nil".

A proper [Error] instance will return false if either the [Error] "Code"
field value is not zero (0), or if the [Error] "Text" field value is not
zero length. Conversely, an [Error] instance bearing only a "Diag" field
value will be regarded as zero when evaluted using [ZeroError], and thus
it will return a value of true.

# Important

Using the standard Go idiomatic expression, or its inverse, upon an
instance of [Error] COULD provide misleading results if only the "Diag"
field is populated, even in a "non error" context. Keep this in mind.
*/
func ExampleZeroError() {

	goErr := ZeroError(exampleGoError) // Any standard Go error
	realErr := ZeroError(exampleError) // Result 66: Operation not allowed on non-leaf
	noErr := ZeroError(exampleNoError) // Result 0:  success
	diag := ZeroError(exampleDiag)     // Result 0:  success, with a diagnostic message

	fmt.Printf("Go Error:   %t\n", goErr)
	fmt.Printf("Real Error: %t\n", realErr)
	fmt.Printf("No Error:   %t\n", noErr)
	fmt.Printf("Diag Only:  %t\n", diag)
	fmt.Printf("Idiomatic:  %t\n", exampleDiag == nil) // inappropriate idiomatic evaluation
	// Output:
	// Go Error:   false
	// Real Error: false
	// No Error:   true
	// Diag Only:  true
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

var exampleError error = LDAPResultNotAllowedOnNonLeaf.New("Operation not allowed on non-leaf")
var exampleNoError error = LDAPResultSuccess.New()
var exampleGoError error = fmt.Errorf("Standard Go error")
var exampleDiag error

func init() {
	diag := LDAPResultSuccess.New().(Error)
	diag.SetDiag("Diagnostic message")
	exampleDiag = diag
}
