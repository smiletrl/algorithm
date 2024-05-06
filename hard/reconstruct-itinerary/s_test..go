package graph

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestFindItinerary(t *testing.T) {
	cases := []struct {
		name    string
		tickets [][]string
		expect  []string
	}{
		{
			name:    "case 1",
			tickets: [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
			expect:  []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
		{
			name:    "case 2",
			tickets: [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}},
			expect:  []string{"JFK", "NRT", "JFK", "KUL"},
		},
		{
			name:    "case 3",
			tickets: [][]string{{"EZE", "AXA"}, {"TIA", "ANU"}, {"ANU", "JFK"}, {"JFK", "ANU"}, {"ANU", "EZE"}, {"TIA", "ANU"}, {"AXA", "TIA"}, {"TIA", "JFK"}, {"ANU", "TIA"}, {"JFK", "TIA"}},
			expect:  []string{"JFK", "ANU", "EZE", "AXA", "TIA", "ANU", "JFK", "TIA", "ANU", "TIA", "JFK"},
		},
		{
			name:    "case 4",
			tickets: [][]string{{"EZE", "TIA"}, {"EZE", "HBA"}, {"AXA", "TIA"}, {"JFK", "AXA"}, {"ANU", "JFK"}, {"ADL", "ANU"}, {"TIA", "AUA"}, {"ANU", "AUA"}, {"ADL", "EZE"}, {"ADL", "EZE"}, {"EZE", "ADL"}, {"AXA", "EZE"}, {"AUA", "AXA"}, {"JFK", "AXA"}, {"AXA", "AUA"}, {"AUA", "ADL"}, {"ANU", "EZE"}, {"TIA", "ADL"}, {"EZE", "ANU"}, {"AUA", "ANU"}},
			expect:  []string{"JFK", "AXA", "AUA", "ADL", "ANU", "AUA", "ANU", "EZE", "ADL", "EZE", "ANU", "JFK", "AXA", "EZE", "TIA", "AUA", "AXA", "TIA", "ADL", "EZE", "HBA"},
		},
	}
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			ar := findItinerary(ca.tickets)
			assert.Equal(t, ca.expect, ar)
		})
	}
}
