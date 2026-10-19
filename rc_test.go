package common

import (
	"testing"
)

func TestResultCode_New(t *testing.T) {
	var err error = ResultCode(5).New("Compare FALSE")
	if code := err.(Error).ResultCode(); code != 5 {
		t.Fatalf("%s failed: expected LDAPResultCompareFalse, got %d",
			t.Name(), code)
	}
}
