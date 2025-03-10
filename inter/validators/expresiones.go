package validators

import "regexp"

var letterRegex *regexp.Regexp
var letterOrNumberRegex *regexp.Regexp
var NumberRegex *regexp.Regexp

func Init() {
	letterRegex = regexp.MustCompile(`/[a-z]/ig`)
	letterOrNumberRegex = regexp.MustCompile(`/[a-z]|[0-9]/ig`)
	NumberRegex = regexp.MustCompile(`[0-9]`)
}

func IsLetter(letter string) bool {
	return letterRegex.Match([]byte(letter))
}

func IsLetterOrNumber(letter string) bool {
	return letterOrNumberRegex.Match([]byte(letter))
}

func IsNumber(num string) bool { return NumberRegex.Match([]byte(num)) }
