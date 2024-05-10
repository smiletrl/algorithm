# solution

This solution is based on original solution from problem [valid-arrangement-of-pairs](https://github.com/smiletrl/algorithm/tree/main/hard/valid-arrangement-of-pairs).

Some improvement:

Graph data structure is

`map[vertex]adjQueue`

adjQueue will be sorted in lexical ascending order, so the vertex will small ascending order walks firstly.

The unused edges, i.e, the remaining out vertices from every walk is saved as a stack list.

Let's say the first walk is

`1 - 2 - 6 - 8 - 1`, and the remaining vertices look like `[1, 2, 6, 1]`. In this case, we use the last vertex `1` from remaining vertices stack, and use it as the next walk's start vertex.

Let's say the second walk is

`1 - 5 - 6 -1`, and we join it with first walk, and we get

`1 - 2 - 6 - 8 - [1 - 5 - 6 -1]`

The virtual edge will be the first walk's starting edge, if virtual edge is added.
