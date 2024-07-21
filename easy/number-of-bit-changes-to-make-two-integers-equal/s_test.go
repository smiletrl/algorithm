package solution

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestMinimumAverage(t *testing.T) {
	cases := []struct {
		name   string
		n      int
		k      int
		expect int
	}{
		{
			name:   "case 1",
			n:      13,
			k:      4,
			expect: 2,
		},
	}

	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			r := minChanges(ca.n, ca.k)
			assert.Equal(t, ca.expect, r)
		})
	}
}
