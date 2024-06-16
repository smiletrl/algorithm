package solution

func countCompleteDayPairs(hours []int) int64 {
	m := make(map[int]int, 24)
	for i := 0; i < 24; i++ {
		m[i] = 0
	}
	for _, hour := range hours {
		m[hour%24]++
	}

	count := int64(0)

	if m[0] > 1 {
		count += int64(m[0] * (m[0] - 1) / 2)
	}

	for i := 1; i < 12; i++ {
		count += int64(m[i] * m[24-i])
	}

	if m[12] > 1 {
		count += int64(m[12] * (m[12] - 1) / 2)
	}

	return count
}
