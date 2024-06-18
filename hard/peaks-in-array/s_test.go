package solution

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCountOfPeaks(t *testing.T) {
	cases := []struct {
		name    string
		nums    []int
		queries [][]int
		expect  []int
	}{
		{
			name: "case 1",
			nums: []int{3, 1, 4, 2, 5},
			queries: [][]int{
				{2, 3, 4},
				{1, 0, 4},
			},
			expect: []int{0},
		},
		{
			name: "case 2",
			nums: []int{4, 1, 4, 2, 1, 5},
			queries: [][]int{
				{2, 2, 4},
				{1, 0, 2},
				{1, 0, 4},
			},
			expect: []int{0, 1},
		},
		{
			name: "case 3",
			nums: []int{5, 3, 8, 4},
			queries: [][]int{
				{2, 0, 2},
				{1, 0, 3},
				{1, 1, 3},
			},
			expect: []int{1, 1},
		},
		{
			name: "case 4",
			nums: []int{4, 9, 4, 10, 7},
			queries: [][]int{
				{2, 3, 2},
				{2, 1, 3},
				{1, 2, 3},
			},
			expect: []int{0},
		},
		{
			name: "case 5",
			nums: []int{3, 6, 9},
			queries: [][]int{
				{1, 1, 1},
				{1, 2, 2},
				{2, 2, 3},
			},
			expect: []int{0, 0},
		},
		{
			name: "case 6",
			nums: []int{4, 10, 8, 6},
			queries: [][]int{
				{1, 0, 3},
				{1, 2, 3},
				{1, 2, 3},
			},
			expect: []int{1, 0, 0},
		},
	}
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			r := countOfPeaks(ca.nums, ca.queries)
			assert.Equal(t, ca.expect, r)
		})
	}
}
