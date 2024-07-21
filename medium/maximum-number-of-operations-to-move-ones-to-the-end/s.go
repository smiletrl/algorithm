package solution

import "fmt"

func maxOperations(s string) int {
	gap := 1
	start := 0
	total := 0
	lastStr := ""
	for i := len(s) - 1; i >= 0; i-- {
		a := fmt.Sprintf("%c", s[i])

		if a == "0" {
			if a == lastStr {
				continue
			}
			if start == 0 {
				start = 1
			} else {
				gap++
			}
		} else {
			if start > 0 {
				total = total + gap
			}
		}
		lastStr = a
	}
	return total
}
