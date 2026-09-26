package iteration

import "strings"

func Repeat(letter string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(letter)
	}
	return repeated.String()
}
