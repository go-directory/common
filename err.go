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
}

/*
Error returns the string error message.
*/
func (r Error) Error() string { return r.Text }

/*
NewError returns an instance of error circumscribing
the concrete [Error] type.
*/
func newRes(rc uint16, msg ...string) (err error) {
	if rc == 0 {
		return
	}

	E := Error{Code: rc}

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
