package solution

import "sort"

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	length := len(rewardValues)
	totalMap := make(map[int]struct{})
	for i := 0; i < length; i++ {
		totalMap[rewardValues[i]] = struct{}{}
		if i == 0 {
			continue
		} else {
			// add current position's value if possible. Previous max reward value is less than current position's value.
			for k := range totalMap {
				if k < rewardValues[i] {
					totalMap[k+rewardValues[i]] = struct{}{}
				}
			}
		}
	}

	max := -1
	for k := range totalMap {
		if k > max {
			max = k
		}
	}
	return max
}
