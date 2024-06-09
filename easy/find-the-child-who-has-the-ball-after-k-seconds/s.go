package solution

func numberOfChild(n int, k int) int {
	rounds := k / (n - 1)
	left := k % (n - 1)
	if rounds%2 == 0 {
		return left
	} else {
		return (n - 1) - left
	}
}
