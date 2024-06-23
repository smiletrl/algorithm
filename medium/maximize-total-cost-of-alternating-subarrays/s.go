package solution

func maximumTotalCost(nums []int) int64 {
	le := len(nums)
	sum := make([][]int64, le)
	// sum[i][0] -> +num[i]
	// sum[i][1] -> -num[i]
	sum[0] = []int64{int64(nums[0]), int64(nums[0])}
	for i := 1; i < le; i++ {
		plus1 := sum[i-1][0] + int64(nums[i])
		plus2 := sum[i-1][1] + int64(nums[i])

		if plus1 < plus2 {
			plus1 = plus2
		}

		sum[i] = []int64{plus1, sum[i-1][0] - int64(nums[i])}
	}

	if sum[le-1][0] > sum[le-1][1] {
		return sum[le-1][0]
	}
	return sum[le-1][1]
}
