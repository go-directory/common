package common

import (
	"strings"
)

/*
Error implements the Go error interface type, and contains
an LDAP result code and (usually) a text message.

Instances of this type are created using [NewError].

A zero instance of this type equates to LDAP result code 0
(success).
*/
type Error struct {
	Code uint16
	Text string
	Diag string
}

/*
Error returns the string error message.
*/
func (r Error) Error() string { return r.Text }

/*
SetDiag assigns the input text sequence as a diagnostic
message.

A diagnostic message does not necessarily indicate an
error has occurred.
*/
func (r *Error) SetDiag(diag ...string) {
	if len(diag) > 0 {
		res := &strings.Builder{}
		for i := 0; i < len(diag); i++ {
			res.WriteString(diag[i])
		}
		r.Diag = res.String()
	}
}

/*
NewError returns an instance of error circumscribing
the concrete [Error] type.
*/
func newRes(rc uint16, msg ...string) (err error) {
	var E Error
	if rc == 0 {
		return E
	}

	E.Code = rc

	if len(msg) > 0 {
		res := &strings.Builder{}
		for i := 0; i < len(msg); i++ {
			res.WriteString(msg[i])
		}
		E.Text = res.String()
	}

	return E
}

/*
ErrorIs returns a Boolean value indicative of both of the
following conditions evaluating as true:

  - err is an instance of [Error]
  - err contains a result code matching rc

If the assertion value rc is zero (0) and err is nil, a
result of true is returned, regardless of the underlying
concrete type.

Any diagnostic message present in the error is not evaluated.
*/
func ErrorIs(err error, rc uint16) (is bool) {
	empty := err == nil
	if is = empty && rc == 0; is {
		return
	}
	if !empty {
		switch tv := err.(type) {
		case Error:
			is = tv.Code == rc
		default:
		}
	}

	return
}

/*
ZeroError returns a Boolean value indicative of the input
error instance being zero, or nil depending on use case.

When an instance of [Error] is provided at input, the "Code"
field MUST be 0 and the "Text" field MUST be "" in order to
guarantee a return value of true.

Note that if an instance of [Error] contains ONLY a "Diag"
field value and is evaluated in standard Go idiomatic terms,
this can return a misleading Boolean.

When the input instance is NOT [Error] -- instead, it is
any standard Go error -- it is evaluated using the standard
idiomatic expression "err == nil".
*/
func ZeroError(err error) (is bool) {
	if E, ok := err.(Error); ok && err != nil {
		is = len(E.Text) == 0 && E.Code == 0
	} else {
		is = err == nil
	}

	return
}
