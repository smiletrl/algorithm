package solution

import "fmt"

func doesAliceWin(s string) bool {
	vowels := map[string]struct{}{
		"a": {},
		"e": {},
		"i": {},
		"o": {},
		"u": {},
	}
	for _, i := range s {
		a := fmt.Sprintf("%c", i)
		if _, ok := vowels[a]; ok {
			return true
		}
	}
	return false
}
