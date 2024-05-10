# solution

This is a classic Eulerian path issue https://en.wikipedia.org/wiki/Eulerian_path. Solution here uses Hierholzer's algorithm

Graph data structure is

map[vertex]adjQueue

Each vertex's value is a queue, which holds the adjacent vertices of this key vertex. When a new edge is added, the edge start vertex is the key value, and the edge end vertex is added to this start vertex's adjQueue.

The input guarantees Eulerian trail already, but not a guaranteed Eulerian circuit. If input is not circuit, add a virtual edge to make it circuit, ie. connect the final vertex back to the start vertex. Then Hierholzer's algorithm can be used.

If we need virtual edge, the virtual edge start vertex will have larger input degree than out degree, and edge end vertex will have smaller input degree than our degree.

As per Hierholzer's algorithm, let's say the first walk, we get circuit

`1 - 2 - 6 - 1`

and vertex 2 has unused out edge `2 - 5`. Then, the second walk might look like

`2 - 5 - 4 - 2`

and we join the second walk back to the first walk, which looks like below now:

`1 - [2 - 5 - 4- 2] - 6 - 1`

Repeat above process until no unused edges are left.

Finally, we cut from the virtual edge to get fianl result, if virtual edge has been added. Let's say the final circuit is

`1 - 2 - 5 - 4 - 2 - 6 - 1`

The virtual edge is `5 - 4`, and we cut it from this virtual edge to get final result:

`4 - 2 - 6 - 1 - 2 - 5`
