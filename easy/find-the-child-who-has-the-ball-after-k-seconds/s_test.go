package solution

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNumberOfChild(t *testing.T) {
	cases := []struct {
		name   string
		n      int
		k      int
		expect int
	}{
		{
			name:   "case 1",
			n:      3,
			k:      5,
			expect: 1,
		},
		{
			name:   "case 2",
			n:      5,
			k:      6,
			expect: 2,
		},
		{
			name:   "case 3",
			n:      4,
			k:      2,
			expect: 2,
		},
	}
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			r := numberOfChild(ca.n, ca.k)
			assert.Equal(t, ca.expect, r)
		})
	}
}
