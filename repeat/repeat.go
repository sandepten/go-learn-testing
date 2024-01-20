package repeat

import "strings"

// Repeat takes a character and repeats it 5 times
func Repeat(character string, times int) string {
	if times == 0 {
		return ""
	}
	return strings.Repeat(character, times)
}
