package graph

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestValidArrangement(t *testing.T) {
	cases := []struct {
		name   string
		edges  [][]int
		expect [][]int
	}{
		{
			name:   "case 1",
			edges:  [][]int{{5, 1}, {4, 5}, {11, 9}, {9, 4}},
			expect: [][]int{{11, 9}, {9, 4}, {4, 5}, {5, 1}},
		},
		{
			name:   "case 2",
			edges:  [][]int{{0, 1}, {2, 0}, {1, 2}},
			expect: [][]int{{1, 2}, {2, 0}, {0, 1}},
		},
		{
			name:   "case 3",
			edges:  [][]int{{5, 3}, {2, 3}, {0, 1}, {1, 4}, {0, 5}, {3, 2}, {4, 3}, {3, 0}},
			expect: [][]int{{0, 5}, {5, 3}, {3, 0}, {0, 1}, {1, 4}, {4, 3}, {3, 2}, {2, 3}},
		},
	}
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			ar := ValidArrangement(ca.edges)
			assert.Equal(t, ca.expect, ar)
		})
	}
}
