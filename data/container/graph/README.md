# Graph data structures
`graph` is a Go package providing building blocks for working with graphs in Go. Most of the types and interfaces are built around `GenericGraph` which is an interface that represents a minimal and practical graph. `GenericGraph` can fit any graphs.

All graphs in gsl implements `GenericGraph`, so, for functions that must work with all graphs, use GenericGraph as input parameter. Otherwise, if your code needs to work with weighted graphs only, then use `GraphWeighted[T, S]`.

For example, for unweighted graphs, I wrote a new interface `GraphWeighted[T, S]` that sits on top of `GenericGraph`. Then I declare a new concrete type `GraphWeightedImpl[T, S]` that implements `GraphWeighted[T, S]`. And if I want to use this `GraphWeightedImpl` for Dijkstra search, I can wrap it with `DijkstraGraph[T, S]`. And finally, if I want to use this graph concurrently, I can wrap `DijkstraGraph[T, S]` with `SafeGraph` again.

In addition to the readily usable `Graph`, `graph` also provides `GenericGraph`, which is very generic and is used by both `Graph` and `wgraph.GraphWeighted`.

Each subpackage in this package provides basic building block and my personal reference implementations for a type of graph. These are the following subpackages:

- `wgraph` for weighted graphs (directional/undirectional weighted graphs). Currently, `wgraph.GraphWeighted` and `Graph` has nothing in common, so they are not interchangable, although `Node` is pretty much unrestricted, so `wgraph.NodeWeighted` can actually be used as a `Node` from this package. See an example code in `wgraph`, where we used both BFS and Dijkstra from the 2 packages on the same data structures.
