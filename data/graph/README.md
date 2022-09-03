# Graph data structures
This package currently provides a `Graph` interface for unweighted graph, and an example type that implements this interface (`GraphImpl`). It also provides other topological sorting methods for `Graph`, such as breadth-first-search.

Each subpackage in this package provides basic building block and my personal reference implementations for a type of graph. These are the following subpackages:

- `wgraph` for weighted graphs (directional/undirectional weighted graphs). Currently, `wgraph.WeightedGraph` and `Graph` has nothing in common, so they are not interchangable, although `Node` is pretty much unrestricted, so `wgraph.WeightedNode` can actually be used as a `Node` from this package. See an example code in `wgraph`, where we used both BFS and Dijkstra from the 2 packages on the same data structures.
