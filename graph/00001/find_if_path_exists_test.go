package graph

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestValidPath(t *testing.T) {
	cases := []struct {
		name        string
		n           int
		edges       [][]int
		source      int
		destination int
		isValid     bool
	}{
		{
			name:        "case 1",
			n:           6,
			edges:       [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {0, 3}},
			source:      0,
			destination: 5,
			isValid:     true,
		},
		{
			name:        "case 2",
			n:           6,
			edges:       [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}},
			source:      0,
			destination: 5,
			isValid:     false,
		},
	}
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			valid := ValidPath(ca.n, ca.edges, ca.source, ca.destination)
			assert.Equal(t, ca.isValid, valid)
		})
	}
}
