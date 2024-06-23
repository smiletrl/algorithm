package solution

func minimumArea(grid [][]int) int {
	xl, xr, yt, yb := 1000, -1, 1000, -1
	for i, row := range grid {
		for j, col := range row {
			if col == 1 {
				if xl > j {
					xl = j
				}
				if xr < j {
					xr = j
				}
				if yt > i {
					yt = i
				}
				if yb < i {
					{
						yb = i
					}
				}
			}
		}
	}
	return (xr - xl + 1) * (yb - yt + 1)
}
