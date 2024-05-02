package graph

// https://leetcode.com/problems/valid-arrangement-of-pairs/description/

func ValidArrangement(pairs [][]int) [][]int {
	g := &Graph{
		adj:                       make(map[int]*adjacentVertices),
		virtualPair:               nil,
		finalCircuit:              nil,
		remainingVertices:         RemainingVertices{},
		walkVertexNode:            nil,
		floatingWalkEndVertexNode: nil,
	}
	for _, pair := range pairs {
		g.addEdge(pair)
		// optional, to get same result. different start vertex will generate different
		// arrangment
		g.startVertex = pair[0]
	}
	g.constructCircuit()
	return g.out()
}

type adjacentVertices struct {
	vertices            []int
	currentIndex        int
	outDegree, inDegree int
}

func (a *adjacentVertices) pop() int {
	if a.currentIndex > len(a.vertices)-1 {
		return -1
	}
	a.currentIndex++

	return a.vertices[a.currentIndex-1]
}

func (a *adjacentVertices) insert(e int) {
	a.vertices = append(a.vertices, e)
}

func (a *adjacentVertices) length() int {
	return len(a.vertices) - a.currentIndex
}

type Graph struct {
	adj map[int]*adjacentVertices

	// used to fix Eulerian circuit from Eulerian trail to Eulerain circuit
	virtualPair []int

	// once a walking circuit is done, join it with the final circuit
	finalCircuit *Circuit

	// any vertices within one walk which still have unused out edges after the walk
	remainingVertices RemainingVertices

	// this vertex is a pop value from remainingVertices
	walkVertexNode *Node

	// when a new walk starts, record this walk start node's next node as the end node for this round walk.
	floatingWalkEndVertexNode *Node

	startVertex int
}

func (g *Graph) addEdge(pair []int) {
	if _, ok := g.adj[pair[0]]; !ok {
		g.adj[pair[0]] = &adjacentVertices{}
	}
	if _, ok := g.adj[pair[1]]; !ok {
		g.adj[pair[1]] = &adjacentVertices{}
	}
	g.adj[pair[0]].insert(pair[1])
	g.adj[pair[0]].outDegree++
	g.adj[pair[1]].inDegree++
}

// find whether we need fix this graph by adding a virtual path to make
// this graph an Eulerian circuit. This graph is assumed to be an Eulerian
// trail already.
func (g *Graph) fixCircuit() {
	start, end := -1, -1
	for vertex, a := range g.adj {
		if a.inDegree < a.outDegree {
			start = vertex
		} else if a.inDegree > a.outDegree {
			end = vertex
		}
	}

	if start != -1 {
		// start and end must be assigned new value together
		// this virtual path will join Eluerian trail's end vertex to start vertex to
		// make it an Eulerian circuit
		g.virtualPair = []int{end, start}
		g.addEdge(g.virtualPair)
	}
}

func (g *Graph) next(v int) int {
	return g.adj[v].pop()
}

// Hierholzer's algorithm
func (g *Graph) walk() {
	// walk from start vertex node
	s := g.walkVertexNode.val

	// record the start node's next node, because start node's next node will be replaced with new node from this
	// round's walk
	g.floatingWalkEndVertexNode = g.walkVertexNode.Next

	// return if current start vertex doesn't have out adj already.
	if g.adj[s].length() == 0 {
		return
	}

	preNode := g.walkVertexNode
	previous_i := s

	i := g.next(s)

	// if edge start vertex still have unused out edges, add it to remaining vertices queue
	if g.adj[previous_i].length() > 0 {
		g.remainingVertices.insert(preNode)
	}

	for i != s {
		// add it to circuit
		preNode = NewNode(i)
		previous_i = i

		g.walkVertexNode.attach(preNode)

		// move walk node to next node
		g.walkVertexNode = preNode

		// get next vertex
		i = g.next(i)

		// if edge start vertex still have unused out edges, add it to remaining vertices queue
		if g.adj[previous_i].length() > 0 {
			g.remainingVertices.insert(preNode)
		}
	}

	// i == s now, i walks back to original vertex. Let's create a new node for original vertex.
	preNode = NewNode(i)
	// the new node will direct to floating walk end vertex node. In this way, we join current walk back to the circuit.
	preNode.Next = g.floatingWalkEndVertexNode
	g.walkVertexNode.attach(preNode)
}

func (g *Graph) initWalk() {
	// find first vertex
	// s := -1
	// for i := range g.adj {
	// 	s = i
	// 	break
	// }
	s := g.startVertex
	startNode := NewNode(s)
	g.finalCircuit = &Circuit{
		start: startNode,
	}
	g.walkVertexNode = startNode
}

func (g *Graph) postWalk() {
	// reset walk vertex node from remaining queue
	g.walkVertexNode = g.remainingVertices.pop()
}

func (g *Graph) constructCircuit() {
	g.fixCircuit()
	g.initWalk()
	for g.walkVertexNode != nil {
		g.walk()
		g.postWalk()
	}
}

func (g *Graph) out() [][]int {
	// get the out pairs
	out := [][]int{}
	it := g.finalCircuit.start
	if g.virtualPair == nil {
		for it.Next != nil {
			out = append(out, []int{it.val, it.Next.val})
			it = it.Next
		}
	} else {
		phase_1 := [][]int{}
		phase_2 := [][]int{}
		phase := 1
		for it.Next != nil {
			// cut from the virtual edge. It's possible that virtual edge has same edges already, and
			// we should only cut circuit once.
			if it.val == g.virtualPair[0] && it.Next.val == g.virtualPair[1] && phase == 1 {
				itNext := it.Next
				it.Next = nil
				it = itNext
				phase = 2
				continue
			}
			if phase == 1 {
				phase_1 = append(phase_1, []int{it.val, it.Next.val})
			} else {
				phase_2 = append(phase_2, []int{it.val, it.Next.val})
			}
			it = it.Next
		}
		out = append(phase_2, phase_1...)
	}
	return out
}

// an Eulerian circuit
type Circuit struct {
	start *Node
}

// node belongs to doubly linked list
type Node struct {
	Previous *Node
	Next     *Node

	// this node's value
	val int
}

func NewNode(val int) *Node {
	return &Node{
		val: val,
	}
}

func (n *Node) attach(an *Node) {
	an.Previous = n
	n.Next = an
}

type RemainingVertices struct {
	vertices     []*Node
	currentIndex int
}

func (r *RemainingVertices) pop() *Node {
	if r.currentIndex > len(r.vertices)-1 {
		return nil
	}
	r.currentIndex++

	return r.vertices[r.currentIndex-1]
}

func (r *RemainingVertices) insert(n *Node) {
	r.vertices = append(r.vertices, n)
}
