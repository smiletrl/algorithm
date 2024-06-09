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
			n:      4,
			k:      5,
			expect: 56,
		},
		{
			name:   "case 2",
			n:      5,
			k:      3,
			expect: 35,
		},
	}
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			r := valueAfterKSeconds(ca.n, ca.k)
			assert.Equal(t, ca.expect, r)
		})
	}
}
