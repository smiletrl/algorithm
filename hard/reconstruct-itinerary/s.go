package graph

import (
	"sort"
)

// https://leetcode.com/problems/reconstruct-itinerary/
// Eularian trail, Eulerian circuit, dfs, queue, stack

func findItinerary(tickets [][]string) []string {
	g := &Graph{
		adj:          make(map[string]*adjacentVertices),
		virtualPair:  nil,
		finalCircuit: nil,
		remainingVertices: RemainingVertices{
			// tickets length is less than 300, so make this
			// stack length to 600.
			vertices:     make([]*Node, 600),
			currentIndex: -1,
		},
		walkVertexNode:            nil,
		floatingWalkEndVertexNode: nil,
		isFirstWalk:               true,
	}
	for _, ticket := range tickets {
		g.addEdge(ticket)
	}

	// set the fixed start vertex
	g.startVertex = "JFK"

	g.constructCircuit()
	return g.out()
}

type adjacentVertices struct {
	// @todo probably add capacity when initialize this queue slice so Go doesn't
	// have to re-assign new memory when insert one new queue item
	vertices            sort.StringSlice
	currentIndex        int
	outDegree, inDegree int
}

func (a *adjacentVertices) pop() string {
	if a.currentIndex > len(a.vertices)-1 {
		return ""
	}
	a.currentIndex++

	return a.vertices[a.currentIndex-1]
}

func (a *adjacentVertices) insert(e string) {
	a.vertices = append(a.vertices, e)
}

func (a *adjacentVertices) length() int {
	return len(a.vertices) - a.currentIndex
}

type Graph struct {
	adj map[string]*adjacentVertices

	// used to fix Eulerian circuit from Eulerian trail to Eulerain circuit
	virtualPair []string

	// once a walking circuit is done, join it with the final circuit
	finalCircuit *Circuit

	// any vertices within one walk which still have unused out edges after the walk
	remainingVertices RemainingVertices

	// this vertex is a pop value from remainingVertices
	walkVertexNode *Node

	// when a new walk starts, record this walk start node's next node as the end node for this round walk.
	floatingWalkEndVertexNode *Node

	// our first walk will start from this vertex
	startVertex string

	// is first walk
	isFirstWalk bool
}

func (g *Graph) addEdge(pair []string) {
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
	start, end := "", ""
	for vertex, a := range g.adj {
		if a.inDegree < a.outDegree {
			start = vertex
		} else if a.inDegree > a.outDegree {
			end = vertex
		}

		// make use of internal Go sort in ascending order for adjacent vertices.
		g.adj[vertex].vertices.Sort()
	}

	if start != "" {
		// start and end must be assigned new value together
		// this virtual path will join Eluerian trail's end vertex to start vertex to
		// make it an Eulerian circuit
		g.virtualPair = []string{end, start}
	}
}

func (g *Graph) next(v string) string {
	return g.adj[v].pop()
}

// Hierholzer's algorithm
func (g *Graph) walk() {
	// walk from start vertex node
	original := g.walkVertexNode
	s := g.walkVertexNode.val

	// record the start node's next node, because start node's next node will be replaced with new node from this
	// round's walk
	g.floatingWalkEndVertexNode = g.walkVertexNode.Next

	// return if current start vertex doesn't have out adj already.
	// if this is first time walk, ignore the length check
	if !g.isFirstWalk && g.adj[s].length() == 0 {
		return
	}

	iterateNode := g.walkVertexNode
	iterate_i := s

	var i string

	// if this is first time walk, walk through virtual edge if it exists
	if g.isFirstWalk && g.virtualPair != nil {
		i = g.virtualPair[1]
	} else {
		i = g.next(s)
	}

	// if edge start vertex still have unused out edges, add it to remaining vertices queue
	if g.adj[iterate_i].length() > 0 {
		g.remainingVertices.insert(iterateNode)
	}

	g.isFirstWalk = false
	for i != s {
		// create a new iterate node and attach it to circuit
		iterateNode = NewNode(i)
		iterate_i = i

		g.walkVertexNode.attach(iterateNode)

		// move walk node to iterate node
		g.walkVertexNode = iterateNode

		// get next iterate vertex
		i = g.next(i)

		// if edge start vertex still have unused out edges, add it to remaining vertices queue
		if g.adj[iterate_i].length() > 0 {
			g.remainingVertices.insert(iterateNode)
		}
	}

	// i == s now, i walks back to original vertex. Let's create a new node for original vertex.
	iterateNode = NewNode(i)
	// the new node will direct to floating walk end vertex node. In this way, we join current walk back to the circuit.
	iterateNode.Next = g.floatingWalkEndVertexNode
	g.walkVertexNode.attach(iterateNode)

	// check it again, so if i has been previously added into remaining queue, use this latter node to
	// process before previous one. In this case, we split the circuit from later position to gain smaller lexical order.
	if g.adj[i].length() > 0 {
		g.remainingVertices.insert(iterateNode)
	}

	its := original
	for its != nil {
		its = its.Next
	}
}

func (g *Graph) initWalk() {
	s := g.startVertex
	// if this graph has virtual pair, we start from the virtual pair, and the first edge to walk
	// is the virtual pair
	if g.virtualPair != nil {
		s = g.virtualPair[0]
	}
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

func (g *Graph) out() []string {
	// get the out pairs
	out := []string{}
	it := g.finalCircuit.start
	if g.virtualPair == nil {
		for it.Next != nil {
			out = append(out, it.val)
			it = it.Next
		}
		// when no virtual pair added, we need include last vertex.
		out = append(out, it.val)
	} else {
		phase_1 := []string{}
		phase_2 := []string{}
		phase := 1
		for it.Next != nil {
			if phase == 1 {
				phase_1 = append(phase_1, it.val)
				// cut from the virtual edge. It's possible that virtual edge has same edges already, and
				// we should only cut circuit once.
				if it.val == g.virtualPair[0] && it.Next.val == g.virtualPair[1] {
					itNext := it.Next
					it.Next = nil
					it = itNext
					phase = 2

					continue
				}
			} else {
				phase_2 = append(phase_2, it.val)
			}
			it = it.Next
		}
		// with virtual pair added, the last vertex is a duplicated vertex as the initial vertex in this circuit,
		// omit it for this itinerary problem.
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
	val string
}

func NewNode(val string) *Node {
	return &Node{
		val: val,
	}
}

func (n *Node) attach(an *Node) {
	an.Previous = n
	n.Next = an
}

type RemainingVertices struct {
	// this will be implemented as stack, ie. FILO.
	// split from circuit's later postion to gain smaller lexical order.
	vertices     []*Node
	currentIndex int
}

func (r *RemainingVertices) pop() *Node {
	if r.currentIndex < 0 {
		return nil
	}

	res := r.vertices[r.currentIndex]
	r.currentIndex--
	return res
}

func (r *RemainingVertices) insert(n *Node) {
	r.currentIndex++
	r.vertices[r.currentIndex] = n
}
