package common

import (
	//"fmt"
	//"strconv"
	"reflect"
	"strings"
)

/*
Kind represents the specific unsigned integer type
selected for use within instances of BitValue.
*/
type Kind uint8

/*
Kind constants define the desired bit allocation
size for an instance of BitValue.
*/
const (
	_      Kind = iota // 0x0
	Uint8              // 0x1; allows eight (8) bits, max val: 255
	Uint16             // 0x2; allows sixteen (16) bits, max val: 65535
	Uint32             // 0x3; allows thirty two (32) bits, max val: 4294967295
)

/*
String returns the string name of the receiver instance,
each of which are literals of the underlying type which
was selected during initialization:

  - uint8
  - uint16
  - uint32
*/
func (r Kind) String() (k string) {
	k = `unknown`

	switch r {
	case Uint8:
		k = `uint8`
	case Uint16:
		k = `uint16`
	case Uint32:
		k = `uint32`
	}

	return
}

/*
Size returns the bit size of the receiver. Possible (valid)
values are eight (8), sixteen (16) or thirty-two (32).

A value of zero (0) indicates the instance has not yet been
initialized (hint: see the New function).
*/
func (r Kind) Size() (size int) {
	size = 0

	switch r {
	case Uint8:
		size = 8
	case Uint16:
		size = 16
	case Uint32:
		size = 32
	}

	return
}

/*
BitValue contains the allocated value type, a Kind and a bit size reflecting the
allocated instance magnitude.

Shift, Unshift and Positive operations may be conducted against instances of this
type.

New instances of this type are created using the New package-level function.
*/
type BitValue struct {
	k Kind           // user-selected Kind
	s uint8          // size (in bits: 8, 16 or 32)
	v any            // allocated instance (as a ptr), per Kind
	m map[int]string // string names for values, optional
}

/*
Value returns the unasserted (POINTER) instance of an
interface (any) within the receiver instance. If the
value is nil, this indicates the receiver instance has
not yet been initialized (hint: see New function).
*/
func (r BitValue) Value() any {
	return r.v
}

/*
NamesMap returns the instance of map[int]string found
within the receiver instance, else nil is returned.

The map is used to resolve string names to shift values
(e.g.: consts).
*/
func (r BitValue) NamesMap() map[int]string {
	return r.m
}

/*
SetNamesMap assigns an instance of map[int]string to
the receiver. The instance, if non-nil, shall be used
to resolve string names to shift values (e.g.: consts).

Case is not significant in the string matching process.
*/
func (r *BitValue) SetNamesMap(m map[int]string) {
	r.m = m
}

/*
Int returns the integer form of the underlying
allocated value. As only unsigned integer types
are supported in this package, the value shall
never be less than zero (0).
*/
func (r BitValue) Int() (i int) {
	switch r.k {
	case Uint8:
		i = int(*(r.v.(*uint8)))
	case Uint16:
		i = int(*(r.v.(*uint16)))
	case Uint32:
		i = int(*(r.v.(*uint32)))
	}

	return
}

/*
Kind returns the instance of Kind assigned to the receiver
instance, which can be one (1) of three (3) possible values.

See the Kind constants for the complete list.
*/
func (r BitValue) Kind() Kind {
	return r.k
}

/*
Shift shall left-shift the bits within the receiver to include
input value(s) x.

If any of values x are considered extremes in their magnitudes
(either zero or the maximum allowed per allocation size), one (1)
of the following shall occur:

  - If value is the maximum, receiver will be clobbered with it (overwritten), i.e.: shift all
  - If value is zero (0), no shift shall occur as it is illogical and not actionable

Additionally, if a maximum value is present within a variadic call
containing >1 slice, unexpected results may ensue. Generally, use
of the maximum value should be used in unary context when calling
this method.
*/
func (r BitValue) Shift(x ...any) BitValue {
	for i := 0; i < len(x); i++ {
		if X, ok := r.verifyShiftValue(x[i]); ok {
			r.shift(X)
		}
	}

	return r
}

/*
None is a convenience method that calls r.Unshift(r.Max()), which will
set the underlying integer value to zero (0) (allocated minimum).
*/
func (r BitValue) None() BitValue {
	r.Unshift(r.Max())
	return r
}

/*
All is a convenience method that calls r.Shift(r.Max()), which
will set the underlying integer value to ^uintN(0) (allocated
maximum).
*/
func (r BitValue) All() BitValue {
	r.Shift(r.Max())
	return r
}

/*
Unshift shall right-shift the bits within the receiver to remove
input value(s) x.

If any of values x are considered extremes in their magnitudes
(either zero or the maximum allowed per allocation size), one (1)
of the following shall occur:

  - If value is the maximum, receiver will be annihilated to zero (0), i.e.: unshift all
  - If value is zero (0), no unshift shall occur as it is illogical and not actionable

Additionally, if a maximum value is present within a variadic call
containing >1 slice, unexpected results may ensue. Generally, use
of the maximum value should be used in unary context when calling
this method.
*/
func (r BitValue) Unshift(x ...any) BitValue {
	for i := 0; i < len(x); i++ {
		if X, ok := r.verifyShiftValue(x[i]); ok {
			r.unshift(X)
		}
	}

	return r
}

/*
Positive returns a Boolean value indicative of whether input value
x's bits are set within the receiver. Negation (!) implies negative,
or 'bit not set'.
*/
func (r BitValue) Positive(x any) (posi bool) {
	if X, ok := r.verifyShiftValue(x); ok {
		posi = r.positive(X)
	}

	return
}

/*
shift is a private method used by BitValue.Shift.
*/
func (r BitValue) shift(x int) {
	if r.isExtreme(x) {
		r.shiftExtremes(x)
		return
	}

	if !r.positive(x) {
		switch r.k {
		case Uint8:
			*(r.v.(*uint8)) |= uint8(x)
		case Uint16:
			*(r.v.(*uint16)) |= uint16(x)
		case Uint32:
			*(r.v.(*uint32)) |= uint32(x)
		}
	}

	return
}

/*
isExtreme returns a Boolean value indicative of whether
integer value x is either zero (0) or the maximum for
the underlying type.
*/
func (r BitValue) isExtreme(x int) bool {
	return x == r.Max() || x == 0
}

/*
shiftExtremes handles either of two extremes: the
shift value is zero (0) or is the maximum for the
underlying type (^uintN(0), where N is a bitsize
of 8, 16 or 32).

Zero (0) is ignored. Max clobbers receiver.

The name may be somewhat misleading; the extreme
value is not shifted, rather it is set literally
and will clobber whatever value was present at
that point.
*/
func (r BitValue) shiftExtremes(x int) {
	if x == r.Max() {
		switch r.k {
		case Uint8:
			*(r.v.(*uint8)) = ^uint8(0)
		case Uint16:
			*(r.v.(*uint16)) = ^uint16(0)
		case Uint32:
			*(r.v.(*uint32)) = ^uint32(0)
		}
	}
}

/*
unshiftExtremes handles the shifting of the maximum
value permitted by the underlying type in negated
context. In other words, unshift max means unshift
everything.

Zero (0) is ignored.
*/
func (r BitValue) unshiftExtremes(x int) {
	if x == r.Max() {
		switch r.k {
		case Uint8:
			*(r.v.(*uint8)) = 0
		case Uint16:
			*(r.v.(*uint16)) = 0
		case Uint32:
			*(r.v.(*uint32)) = 0
		}
	}
}

/*
unshift is a private method used by BitValue.Unshift.
*/
func (r BitValue) unshift(x int) {
	if r.isExtreme(x) {
		r.unshiftExtremes(x)
		return
	}

	if r.positive(x) {
		switch r.k {
		case Uint8:
			*(r.v.(*uint8)) = *(r.v.(*uint8)) &^ uint8(x)
		case Uint16:
			*(r.v.(*uint16)) = *(r.v.(*uint16)) &^ uint16(x)
		case Uint32:
			*(r.v.(*uint32)) = *(r.v.(*uint32)) &^ uint32(x)
		}
	}

	return
}

/*
positive is a private method used by BitValue.Positive.
*/
func (r BitValue) positive(x int) (posi bool) {
	switch r.k {
	case Uint8:
		posi = *(r.v.(*uint8))&uint8(x) > 0
	case Uint16:
		posi = *(r.v.(*uint16))&uint16(x) > 0
	case Uint32:
		posi = *(r.v.(*uint32))&uint32(x) > 0
	}

	return
}

/*
toInt is merely a means to keep cyclomatics low in this package. this
package-level function simply asserts and casts various integer types
to a straight int instance. If value x was an int to begin with, it is
silently accepted and treated as if it were something else.
*/
func toInt(x any) (v int, ok bool) {
	switch tv := x.(type) {
	case int:
		v = tv
		ok = true
	case uint8:
		v = int(tv)
		ok = true
	case uint16:
		v = int(tv)
		ok = true
	case uint32:
		v = int(tv)
		ok = true
	default:
		// for custom user-defined types based on
		// select built-in integer types.
		const maxU = ^uint(0)
		const minU = 0
		const maxI = int(maxU >> 1)
		const minI = -maxI - 1

		rfv := reflect.ValueOf(tv)
		switch k := rfv.Kind(); k {
		case reflect.Int8, reflect.Int16, reflect.Int32:
			if i := rfv.Int(); int64(minI) <= int64(i) && int64(i) <= int64(maxI) {
				v = int(i)
				ok = true
			}
		case reflect.Uint8, reflect.Uint16, reflect.Uint32:
			if u := rfv.Uint(); int64(u) <= int64(maxI) {
				v = int(u)
				ok = true
			}
		}
	}

	return
}

/*
Max returns ^uintN(0), where N is a bit size of eight
(8), sixteen (16), or thirty-two (32).

The returned integer represents the largest possible
value permitted by the underlying allocated instance.
*/
func (r BitValue) Max() (max int) {
	switch r.k {
	case Uint8:
		max = int(^uint8(0))
	case Uint16:
		max = int(^uint16(0))
	case Uint32:
		max = int(^uint32(0))
	}

	return
}

/*
Size returns the underlying bit size of the receiver.
*/
func (r BitValue) Size() (size int) {
	return r.k.Size()
}

/*
Min returns zero (0), which indicates the lowest permitted
value within the receiver for any supported allocation.
*/
func (r BitValue) Min() (min int) {
	return 0
}

/*
verifyShiftValue returns an integer and a Boolean value following
the processing of input value x, which must represent some kind of
integer. Assuming that integer is not out-of-bounds (in terms of
minimum or maximum value magnitudes permitted) for the underlying
type instance, it is returned alongside a success-indicative Boolean
value.

This method also resolves a string name to a known int bit value,
if set within the NamesMap within the receiver.
*/
func (r BitValue) verifyShiftValue(x any) (X int, ok bool) {
	// if the input value was a string, try
	// to resolve it to an int value ...
	if str, asserted := x.(string); asserted {
		x = r.strIndex(str)
	}

	if X, ok = toInt(x); ok {
		ok = r.Min() <= X && X <= r.Max()
	}

	return
}

/*
strIndex returns the integer index for the specified bitvalue,
if found within the underlying names map. Otherwise, -1 is
returned.
*/
func (r BitValue) strIndex(x string) (idx int) {
	idx = -1

	for k, v := range r.m {
		if strings.EqualFold(v, x) {
			idx = k
		}
	}

	return
}

/*
New initializes a new instance of [BitValue], using [Kind] k as
the indicator for the desired bit allocation size. See the [Kind]
constants for available values.
*/
func New(k Kind) (bv BitValue) {
	bv.k = k

	switch k {
	case Uint8:
		bv.s = uint8(k.Size())
		bv.v = new(uint8)

	case Uint16:
		bv.s = uint8(k.Size())
		bv.v = new(uint16)

	case Uint32:
		bv.s = uint8(k.Size())
		bv.v = new(uint32)
	}

	return
}
