package solution

func countOfPeaks(nums []int, queries [][]int) []int {
	le := len(nums)
	res := []int{}
	peaks := make([]int, le)
	peaks[0] = 0
	peaks[le-1] = 0
	b := bit{
		le:   le,
		tree: make([]int, le+1),
	}

	var i int
	for i = 1; i < le-1; i++ {
		if (nums[i] > nums[i-1]) && (nums[i] > nums[i+1]) {
			peaks[i] = 1
		}
	}

	for i = 0; i < le; i++ {
		b.update(i, peaks[i])
	}

	for _, query := range queries {
		ty := query[0]
		switch ty {
		case 2:
			nums[query[1]] = query[2]

			// re-calculate peaks for index query[1], query[1] - 1, query[1] + 1
			for i := query[1] - 1; i <= query[1]+1; i++ {
				if i <= 0 {
					continue
				}
				if i >= le-1 {
					continue
				}
				old := peaks[i]
				if (nums[i] > nums[i-1]) && (nums[i] > nums[i+1]) {
					peaks[i] = 1
				} else {
					peaks[i] = 0
				}

				delta := peaks[i] - old
				if delta != 0 {
					b.update(i, delta)
				}
			}

		case 1:
			if query[2] == query[1] {
				res = append(res, 0)
			} else {
				res = append(res, b.sum(query[2]-1)-b.sum(query[1]))
			}
		}
	}
	return res
}

// binary tree
type bit struct {
	tree []int
	le   int
}

func (b *bit) update(i int, delta int) {
	i++
	for ; i <= b.le; i += (i & -i) {
		b.tree[i] += delta
	}
}

func (b *bit) sum(i int) int {
	count := 0
	i++
	for ; i > 0; i -= (i & -i) {
		count += b.tree[i]
	}
	return count
}
