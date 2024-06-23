package solution

import (
	"sort"
)

func minimumAverage(nums []int) float64 {
	n := len(nums)
	sort.Ints(nums)
	sav := float64(nums[n-1]) + float64(1)
	j := 0
	for i := 0; i < n/2; i++ {
		av := float64((float64(nums[j]) + float64(nums[n-1-j])) / 2)
		if sav > av {
			sav = av
		}
		j++
	}
	return sav
}
