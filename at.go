package common

/*
isAttributeDescriptor scans the input []byte val and judges
whether it appears to qualify as a valid RFC 4512 descriptor
(or "descr"), in that:

  - it begins with an alpha
  - it ends with an alpha or digit
  - it contains only alphas, digits, hyphens or semicolons
  - it contains no consecutive hyphens or semicolons

Attribute tags, such as ";lang-jp" or ";binary", are tolerated.
*/
func IsAttributeDescriptor(val []byte) bool {
	if len(val) == 0 {
		return false
	}

	// must begin with an alpha.
	if !isAlpha(rune(val[0])) {
		return false
	}

	// can only end in alnum.
	if !isAlnum(rune(val[len(val)-1])) {
		return false
	}

	for i := 0; i < len(val); i++ {
		ch := rune(val[i])
		switch {
		case isAlnum(ch):
			// ok
		case ch == ';', ch == '-':
			// ok
		default:
			return false
		}
	}

	return true
}
