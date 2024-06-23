package solution

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMaximumTotalCost(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		expect int64
	}{
		{
			name:   "case 1",
			nums:   []int{1, -2, 3, 4},
			expect: 10,
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			r := maximumTotalCost(ca.nums)
			assert.Equal(t, ca.expect, r)
		})
	}
}
