package solution

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLongestCommonSubpath(t *testing.T) {
	cases := []struct {
		name        string
		n           int
		paths       [][]int
		expectedLen int
	}{
		{
			name: "case -1",
			n:    5,
			paths: [][]int{
				{0, 1, 2, 3, 4},
				{2, 3, 4},
				{4, 0, 1, 2, 3},
			},
			expectedLen: 2,
		},
		{
			name: "case 0",
			n:    4,
			paths: [][]int{
				{1, 2, 3},
				{2},
			},
			expectedLen: 1,
		},
		{
			name: "case 1",
			n:    5,
			paths: [][]int{
				{1, 2, 3},
				{2, 3, 4},
			},
			expectedLen: 2,
		},
		{
			name: "case 2",
			n:    10,
			paths: [][]int{
				{1, 2, 3, 8, 9},
				{2, 3, 8, 7, 7},
				{12, 8, 2, 3, 8},
			},
			expectedLen: 3,
		},
		{
			name: "case 3",
			n:    5,
			paths: [][]int{
				{0, 1, 2, 3, 4},
				{4, 3, 2, 1, 0},
			},
			expectedLen: 1,
		},
	}
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			sub := longestCommonSubpath(ca.n, ca.paths)
			assert.Equal(t, ca.expectedLen, sub)
		})
	}
}
