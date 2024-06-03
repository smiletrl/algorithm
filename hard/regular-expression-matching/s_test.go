package solution

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestIsmatched(t *testing.T) {
	cases := []struct {
		name      string
		s         string
		p         string
		isMatched bool
	}{
		{
			name:      "case 1",
			s:         "aa",
			p:         "a",
			isMatched: false,
		},
		{
			name:      "case 2",
			s:         "aa",
			p:         "a*",
			isMatched: true,
		},
		{
			name:      "case 3",
			s:         "aab",
			p:         "c*a*b",
			isMatched: true,
		},
		{
			name:      "case 4",
			s:         "ab",
			p:         ".*",
			isMatched: true,
		},
		{
			name:      "case 5",
			s:         "mississippi",
			p:         "mis*is*ip*.",
			isMatched: true,
		},
		{
			name:      "case 6",
			s:         "aaa",
			p:         "aaaa",
			isMatched: false,
		},
	}
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			matched := isMatch(ca.s, ca.p)
			assert.Equal(t, ca.isMatched, matched)
		})
	}
}

func TestLoop(t *testing.T) {
	loopStr("sjskdl")
}
