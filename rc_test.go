package common

import (
	"testing"
)

func TestResultCode_New(t *testing.T) {
	err := ResultCode(5).New("Compare FALSE")
	if code := err.(Error).Code; code != 5 {
		t.Fatalf("%s failed: expected LDAPResultCompareFalse, got %d",
			t.Name(), code)
	}
}
