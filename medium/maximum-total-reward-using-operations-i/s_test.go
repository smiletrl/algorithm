package solution

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMaxTotalReward(t *testing.T) {
	cases := []struct {
		name   string
		i      []int
		expect int
	}{
		{
			name:   "case 1",
			i:      []int{1, 1, 3, 3},
			expect: 4,
		},
		{
			name:   "case 2",
			i:      []int{1, 6, 4, 3, 2},
			expect: 11,
		},
	}
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			r := maxTotalReward(ca.i)
			assert.Equal(t, ca.expect, r)
		})
	}
}
