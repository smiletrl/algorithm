package solution

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCountCompleteDayPairs(t *testing.T) {
	cases := []struct {
		name   string
		hours  []int
		expect int64
	}{
		{
			name:   "case 1",
			hours:  []int{12, 12, 30, 24, 24},
			expect: 2,
		},
		{
			name:   "case 2",
			hours:  []int{72, 48, 24, 3},
			expect: 3,
		},
		{
			name:   "case 3",
			hours:  []int{11, 22, 2, 25, 19, 2},
			expect: 2,
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			r := countCompleteDayPairs(ca.hours)
			assert.Equal(t, ca.expect, r)
		})
	}
}
