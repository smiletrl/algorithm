package solution

import (
	"sort"
)

func maximumTotalDamage(power []int) int64 {
	sort.Ints(power)

	max := int64(0)

	// map of each p's total value. total value can be smaller than or equal to p's max value.
	m := map[int]int64{}

	// map of each p's max value. This value is sorted in ascending order.
	maxp := make(map[int]int64)

	for i, p := range power {
		if i == 0 {
			m[p] = int64(p)
		} else if p == power[i-1] {
			m[p] += int64(p)
		} else {
			pre := int64(0)
			for j := i - 1; j > -1; j-- {
				if power[j] <= p-3 {
					pre = maxp[power[j]]
					break
				}
			}
			m[p] = pre + int64(p)
		}

		if m[p] > max {
			max = m[p]
		}

		// log this p's max value
		maxp[p] = max
	}

	return max
}
