package common

import (
	"bytes"
	"strconv"
	"strings"
)

var (
	sLC = strings.ToLower
)

type strBuilder = strings.Builder

var (
	bLC  = bytes.ToLower
	bEQ  = bytes.Equal
	bEQF = bytes.EqualFold
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
