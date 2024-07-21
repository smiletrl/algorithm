package solution

func minChanges(n int, k int) int {
	if n == k {
		return 0
	}
	if n < k {
		return -1
	}

	c := n | k
	if c != n {
		return -1
	}

	a := n ^ k
	d := 0
	for {
		if a == 0 {
			break
		}
		if a%2 > 0 {
			d++
		}
		a = a >> 1
	}

	return d
}
