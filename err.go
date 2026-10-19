package common

import (
	"errors"
	"strings"
)

/*
Error implements the combination of the standard Go error
with a so-called "LDAPResult Precursor".

Instances of this type are created using [NewError].

Instances of this type implement the Go error interface like
any other error. However the following RFC4511 components are
accessible for getting and setting:

  - Result code (uint16)
  - Matched DN (string)
  - Diagnostic Message (string)
  - Referral ([]string)

Additionally, other helpful components (not standards-derived)
are available:

  - Error "header" (classification label)
  - Error slices (for when multiple errors may be accumulated for a single "event")

A zero instance of this type equates to LDAP result code 0
(success).
*/
type Error struct {
	*ldapResult
}

type ldapResult struct {
	ResultCode        uint16
	Header            string
	Text              string
	DiagnosticMessage string
	MatchedDN         string
	Referral          []string
	Errors            []error
}

/*
Error returns the string error message specified as
the "Text" component.

See also the [Error.Text] and [Error.SetText] methods.
*/
func (r Error) Error() string {
	var e string

	if !r.IsNil() && r.ldapResult.Text != "" {
		if r.ldapResult.Header != "" {
			e = r.ldapResult.Header + ` `
		}
		e += r.ldapResult.Text
	}

	return e
}

/*
AppendError appends slices of error instances to the
receiver instance. Slices are evaluated by way of the
[ZeroError] function.
*/
func (r Error) AppendError(err ...error) {
	if r.IsNil() {
		r.ldapResult = &ldapResult{}
	}

	for i := 0; i < len(err); i++ {
		if !ZeroError(err[i]) {
			r.ldapResult.Errors =
				append(r.ldapResult.Errors, err[i])
		}
	}
}

/*
Errors returns slices of errors.
*/
func (r Error) Errors() []error {
	var errs []error
	if !r.IsNil() {
		errs = r.ldapResult.Errors
	}

	return errs
}

/*
GoError returns the receiver instance as a standard
Go error (using [errors.New]).

Note that this is a "lossy" method, in that only the
"Text" component of [Error] is incorporated into the
return instance. All other components -- such as the
"ResultCode" or "DiagnosticMessage" components -- are
not transferred.

If no "Text" value was set within the receiver instance,
the return value is guaranteed to be nil.
*/
func (r Error) GoError() error {
	var err error
	if !r.IsNil() && len(r.ldapResult.Text) > 0 {
		err = errors.New(r.ldapResult.Text)
	}

	return err
}

/*
IsNil returns a Boolean value indicative of the embedded
receiver instance value is nil.
*/
func (r Error) IsNil() bool { return r.ldapResult == nil }

/*
SetResultCode assigns the input uint16 value to the receiver
instance as the "ResultCode" component.
*/
func (r Error) SetResultCode(rc uint16) {
	if r.IsNil() {
		r.ldapResult = &ldapResult{}
	}
	r.ldapResult.ResultCode = rc
}

/*
ResultCode returns the underlying "ResultCode" uint16
component value.
*/
func (r Error) ResultCode() uint16 {
	var rc uint16
	if !r.IsNil() {
		rc = r.ldapResult.ResultCode
	}

	return rc
}

/*
SetText assigns the input string message slices to the
receiver instance as the "Text" component. It is this
value that will appear in error text when raised.

See the [Error.Error] method for a means of retrieving
the underlying "Text" component value.

Note that this method clobbers any existing text value
present. See also the [Error.AppendText] method for a
means of "building" a single error message gradually.
*/
func (r Error) SetText(msg ...string) {
	if len(msg) > 0 {
		if r.IsNil() {
			r.ldapResult = &ldapResult{}
		}
		r.ldapResult.Text = makeErrorText(msg...)
	}
}

/*
AppendText appends the input text slices to the underlying
"Text" component. This method is useful when a single error
message is being built gradually.

See also the [Error.SetText] method for a means of overwriting
("clobbering") any preexisting "Text" component value in one
shot.
*/
func (r Error) AppendText(msg ...string) {
	if len(msg) > 0 {
		if r.IsNil() {
			r.ldapResult = &ldapResult{}
		}
		r.ldapResult.Text += makeErrorText(msg...)
	}
}

/*
SetDiagnosticMessage assigns the input text sequence as an
LDAPResult "diagnostic message".

A diagnostic message does not necessarily indicate an
error has occurred.
*/
func (r Error) SetDiagnosticMessage(diag ...string) {
	if len(diag) > 0 {
		if r.IsNil() {
			r.ldapResult = &ldapResult{}
		}
		r.ldapResult.DiagnosticMessage = makeErrorText(diag...)
	}
}

/*
SetHeader assigns the input head slices to the receiver
instance as the leading "Header" component value.

A header will be visible when the receiver instance is
represented as a standard error or when the [Error.Error]
method is called. The header is "prepended" to the text.
*/
func (r Error) SetHeader(head ...string) {
	if len(head) > 0 {
		if r.IsNil() {
			r.ldapResult = &ldapResult{}
		}
		r.ldapResult.Header = makeErrorText(head...)
	}
}

func makeErrorText(msg ...string) (s string) {
	if len(msg) > 0 {
		res := &strings.Builder{}
		for i := 0; i < len(msg); i++ {
			res.WriteString(msg[i])
		}

		s = res.String()
	}

	return s
}

/*
DiagnosticMessage returns the underlying "DiagnosticMessage"
component value, if set.
*/
func (r Error) DiagnosticMessage() string {
	var dm string
	if !r.IsNil() {
		dm = r.ldapResult.DiagnosticMessage
	}

	return dm
}

/*
Header returns the underlying "Header" component value, if
set.
*/
func (r Error) Header() string {
	var head string
	if !r.IsNil() {
		head = r.ldapResult.Header
	}

	return head
}

/*
SetMatchedDN assigns the input string distinguished name to
the receiver instance as the "MatchedDN" component.
*/
func (r Error) SetMatchedDN(matchedDN string) {
	if len(matchedDN) > 0 {
		if r.IsNil() {
			r.ldapResult = &ldapResult{}
		}
		r.ldapResult.MatchedDN = matchedDN
	}
}

/*
MatchedDN returns the underlying "MatchedDN" component value,
if set.
*/
func (r Error) MatchedDN() string {
	var dn string
	if !r.IsNil() {
		dn = r.ldapResult.MatchedDN
	}

	return dn
}

/*
SetReferral appends the input referral string slices to
the receiver instance as the "Referral" component set.
*/
func (r Error) SetReferral(referral ...string) {
	if len(referral) > 0 {
		if r.IsNil() {
			r.ldapResult = &ldapResult{}
		}
		r.ldapResult.Referral =
			append(r.ldapResult.Referral,
				referral...)
	}
}

/*
Referral returns the underlying "Referral" slice component
value, if set.
*/
func (r Error) Referral() []string {
	var ref []string
	if !r.IsNil() {
		ref = r.ldapResult.Referral
	}

	return ref
}

/*
NewError returns a concrete instance of [Error] (not error).

This function is used for managing an [Error] instance's
contents prior to return as a standard Go error.
*/
func NewError(rc uint16, msg ...string) Error {
	return newRes(rc, msg...)
}

func newRes(rc uint16, msg ...string) Error {
	return Error{&ldapResult{
		ResultCode: rc,
		Text:       makeErrorText(msg...),
	}}
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
			is = tv.ldapResult != nil && tv.ldapResult.ResultCode == rc
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
	z := uint16(0)
	if E, ok := err.(Error); ok && err != nil {
		is = len(E.Text) == 0 && E.ResultCode() == z
	} else if E2, ok := err.(Error); ok && err != nil {
		is = len(E2.Text) == 0 && E2.ResultCode() == z
	} else {
		is = err == nil
	}

	return
}
