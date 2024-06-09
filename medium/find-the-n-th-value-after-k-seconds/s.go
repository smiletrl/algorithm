package solution

func valueAfterKSeconds(n int, k int) int {
	var p int = 1e9 + 7
	var d [][]int = make([][]int, k+1)
	d[0] = make([]int, n)
	for i := 0; i < n; i++ {
		d[0][i] = 1
	}
	for i := 1; i <= k; i++ {
		d[i] = make([]int, n)
		d[i][0] = 1
		for j := 1; j < n; j++ {
			d[i][j] = (d[i-1][j] + d[i][j-1]) % p
		}
	}
	return d[k][n-1]
}
