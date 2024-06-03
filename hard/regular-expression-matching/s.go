package solution

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	sArr := loopStr(s)
	pArr := loopStr(p)

	i := 0

	nfa := &NFA{}

	nfa.init(pArr)

	// we will find all possible NFA states from starting letter of pattern string.
	nfa.g.dfsArr([]int{0})

	for i = 0; i < len(sArr); i++ {

		// loop through all marked states for current search char, and see if we find a match
		match := []int{}
		for key := range nfa.g.marked {
			if key < len(pArr) {
				if sArr[i] == pArr[key] || pArr[key] == "." {
					// we will compare against next char of pattern string for next char of search string
					// if key == len(pArr) - 1, then key+1 is the complete NFA state
					match = append(match, key+1)
				}
			}
		}

		// if no match found, break
		if len(match) == 0 {
			return false
		}

		// find all NFA states for next char from search text
		nfa.g.marked = make(map[int]struct{}, 0)
		nfa.g.dfsArr(match)
	}

	for m := range nfa.g.marked {
		// reach the complete NFA state
		if m == len(pArr) {
			return true
		}
	}
	return false
}

func loopStr(s string) []string {
	res := make([]string, len(s))
	for pos, a := range s {
		res[pos] = fmt.Sprintf("%c", a)
	}
	return res
}

type NFA struct {
	g *Graph
}

func (n *NFA) init(pArr []string) {
	n.g = &Graph{
		paths:  make(map[int][]int),
		marked: make(map[int]struct{}),
	}
	for i, ps := range pArr {
		if ps == "*" {
			n.g.addEdge(i, i-1)
			n.g.addEdge(i-1, i)
			n.g.addEdge(i, i+1)
		}
	}
}

type Graph struct {
	paths  map[int][]int
	marked map[int]struct{}
}

// Find all the reachable vertices from source s
func (g *Graph) dfs(s int) {
	g.marked[s] = struct{}{}
	edges, ok := g.paths[s]
	if !ok {
		return
	}

	for _, e := range edges {
		if _, ok := g.marked[e]; !ok {
			g.dfs(e)
		}
	}
}

func (g *Graph) dfsArr(sArr []int) {
	for _, s := range sArr {
		g.dfs(s)
	}
}

func (g *Graph) addEdge(s, e int) {
	if _, ok := g.paths[s]; !ok {
		g.paths[s] = make([]int, 0)
	}
	g.paths[s] = append(g.paths[s], e)
}
