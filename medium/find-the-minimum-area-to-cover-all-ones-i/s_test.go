package solution

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMinimumArea(t *testing.T) {
	cases := []struct {
		name   string
		grid   [][]int
		expect int
	}{
		{
			name: "case 1",
			grid: [][]int{
				{0, 1, 0},
				{1, 0, 1},
			},
			expect: 6,
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			r := minimumArea(ca.grid)
			assert.Equal(t, ca.expect, r)
		})
	}
}
