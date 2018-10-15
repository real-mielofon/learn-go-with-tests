package iteration

import "strings"

// Repeat return 5 charecter
func Repeat(character string, count int) (repeated string) {
	return strings.Repeat(character, count)
}
