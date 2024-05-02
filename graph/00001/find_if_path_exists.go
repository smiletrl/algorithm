package graph

// https://leetcode.com/problems/find-if-path-exists-in-graph/description/

func ValidPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}
	g := &Graph{
		graphs:           make([][]int, n),
		marked:           make(map[int]bool, n),
		finaldestination: destination,
	}
	for _, edge := range edges {
		g.addEdge(edge)
	}

	// @todo bfs might be faster
	g.dfs(source)
	if _, ok := g.marked[destination]; ok {
		return true
	}

	return false
}

type Graph struct {
	graphs           [][]int
	marked           map[int]bool
	finaldestination int
}

func (g *Graph) addEdge(edge []int) {
	g.graphs[edge[0]] = append(g.graphs[edge[0]], edge[1])
	g.graphs[edge[1]] = append(g.graphs[edge[1]], edge[0])
}

func (g *Graph) dfs(source int) {
	for _, v := range g.graphs[source] {
		// if v has marked, skip
		if _, ok := g.marked[v]; ok {
			continue
		}
		g.marked[v] = true
		if v == g.finaldestination {
			return
		} else {
			g.dfs(v)
		}
	}
}
