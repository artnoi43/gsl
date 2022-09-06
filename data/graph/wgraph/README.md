# Weighted Graphs
This package provides building blocks for graph with weighted edges. The graph types are:

1. `GraphWeighted[T, S]`: weighted graph interface, which is constrained `GenericGraph`.

2. `GraphWeightedImpl[T, S]`: default implementations of `GraphWeighted[T, S]`.

3. `DijkstraGraph[T, S]`: a wrapper for `GraphWeightImpl[T, S]`. It checks whether the edge value is negative, and errs if so.

The Dijkstra algorithm here should work with any graph that implement the `GraphWeighted` interfaces in `interface.go`. When using code `dijkstra.go` from, be sure to use `*DijkstraGraphUndirected[T, S]).AddDijkstraEdge` instead of `WeigtedGraph.AddEdge` when adding edges to avoid adding edges with bad weight.