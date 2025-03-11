package validators

func IsDelim(d rune) bool {
	switch d {
	case ' ':
		return true
	case '\n':
		return true
	case '\t':
		return true
	case '\r':
		return true
	default:
		return false
	}
}
