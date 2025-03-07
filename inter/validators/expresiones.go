package validators

import "regexp"

var letterRegex *regexp.Regexp
var letterOrNumberRegex *regexp.Regexp

func Init() {
	letterRegex = regexp.MustCompile(`/[a-z]/ig`)
	letterOrNumberRegex = regexp.MustCompile(`/[a-z]|[0-9]/ig`)
}

func IsLetter(letter string) bool {
	return letterRegex.Match([]byte(letter))
}

func IsLetterOrNumber(letter string) bool {
	return letterOrNumberRegex.Match([]byte(letter))
}
