package graph

var (
	graphs           [][]int
	marked           map[int]bool
	finaldestination int
)

func ValidPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}
	graphs = make([][]int, n)
	for _, edge := range edges {
		graphs[edge[0]] = append(graphs[edge[0]], edge[1])
		graphs[edge[1]] = append(graphs[edge[1]], edge[0])
	}

	marked = make(map[int]bool, n)
	marked[source] = true
	finaldestination = destination
	dfs(source)
	if _, ok := marked[destination]; ok {
		return true
	}

	return false
}

func dfs(source int) {
	for _, v := range graphs[source] {
		// if v has marked, skip
		if _, ok := marked[v]; ok {
			continue
		}
		marked[v] = true
		if v == finaldestination {
			return
		} else {
			dfs(v)
		}
	}
}
