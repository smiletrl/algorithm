package solution

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMaximumTotalDamage(t *testing.T) {
	cases := []struct {
		name   string
		power  []int
		expect int64
	}{
		// {
		// 	name:   "case 1",
		// 	hours:  []int{12, 12, 30, 24, 24},
		// 	expect: 2,
		// },
		{
			name:   "case 2",
			power:  []int{5, 9, 2, 10, 2, 7, 10, 9, 3, 8},
			expect: 31,
		},
		// {
		// 	name:   "case 3",
		// 	power:  []int{7, 1, 6, 3},
		// 	expect: 10,
		// },
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			r := maximumTotalDamage(ca.power)
			assert.Equal(t, ca.expect, r)
		})
	}
}
