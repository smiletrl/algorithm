package solution

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMaxTotalReward(t *testing.T) {
	cases := []struct {
		name   string
		s      string
		expect bool
	}{
		{
			name:   "case 1",
			s:      "leetcoder",
			expect: true,
		},
	}
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			r := doesAliceWin(ca.s)
			assert.Equal(t, ca.expect, r)
		})
	}
}
