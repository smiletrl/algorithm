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
			s:         "cb",
			p:         "?a",
			isMatched: false,
		},
		{
			name:      "case 3",
			s:         "aa",
			p:         "*",
			isMatched: true,
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
