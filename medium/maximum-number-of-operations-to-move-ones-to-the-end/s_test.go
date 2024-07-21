package solution

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMaxTotalReward(t *testing.T) {
	cases := []struct {
		name   string
		s      string
		expect int
	}{
		{
			name:   "case 1",
			s:      "1001101",
			expect: 4,
		},
	}
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			r := maxOperations(ca.s)
			assert.Equal(t, ca.expect, r)
		})
	}
}
