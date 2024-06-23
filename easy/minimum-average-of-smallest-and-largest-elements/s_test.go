package solution

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMinimumAverage(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		expect float64
	}{
		{
			name:   "case 1",
			nums:   []int{7, 8, 3, 4, 15, 13, 4, 1},
			expect: 5.5,
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			r := minimumAverage(ca.nums)
			assert.Equal(t, ca.expect, r)
		})
	}
}
