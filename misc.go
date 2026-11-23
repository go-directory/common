package common

import (
	"bytes"
	"strconv"
	"strings"
	"time"
)

type strBuilder = strings.Builder

var (
	sLC = strings.ToLower
)

var (
	bLC   = bytes.ToLower
	bEQ   = bytes.Equal
	bEQF  = bytes.EqualFold
	bHas  = bytes.Contains
	bHSfx = bytes.HasSuffix
)

var (
	itoa  = strconv.Itoa
	atoi  = strconv.Atoi
	fuint = strconv.FormatUint
	fint  = strconv.FormatInt
	fbool = strconv.FormatBool
	puint = strconv.ParseUint
	pint  = strconv.ParseInt
	pbool = strconv.ParseBool
)

/*
TimeToGenTime marshals the input time.Time t and returns
the formatted generalized time (YYYYMMDDHHmmssZ).

The variadic input argument, hires, will invoke so-called
high-resolution time (YYYYMMDDhhmmss.000000Z).
*/
func TimeToGenTime(t time.Time, hires ...bool) []byte {
	var T []byte

	if len(hires) > 0 && hires[0] {
		T = []byte(t.UTC().Format(`20060102150405.000000`) + `Z`)
	} else {
		T = []byte(t.UTC().Format(`20060102150405`) + `Z`)
	}

	return T
}

/*
GenTimeToTime returns an instance of [time.Time] alongside an
error following an attempt to parse the input []byte timestamp
t.
*/
func GenTimeToTime(t []byte) (time.Time, error) {
	var format string = `20060102150405Z`

	if bHas(t, []byte(`.`)) {
		format = `20060102150405.000000Z`
	}

	return time.Parse(format, string(t))
}

/*
GenTimeNow returns the []byte representation of the current
time as generalized UTC time value.

The variadic input argument, hires, will invoke so-called
high-resolution time (YYYYMMDDhhmmss.000000Z).
*/
func GenTimeNow(hires ...bool) []byte {
	return TimeToGenTime(time.Now().UTC(), hires...)
}
