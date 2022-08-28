# Weighted Graphs
This package provides basic building blocks and example implementations of weighted graphs. Package interfaces include 

I'll be adding more types and algorithms as soon as I learn the new ones.

- `ugraph` undirected, weighted graph implementing interface `UndirectedGraph`

- `unode` a node of a `ugraph`, implements `UndirectedNode`

- `uedge` an edge of a `ugraph`, implements `UndirectedEdge`

> The Dijkstra algorithm here should work with any graphs, nodes, and edges that implement the interfaces in `interface.go`. When using code `dijkstra.go` from, be sure to use `*DijkstraGraphUndirected[T, S]) AddDijkstraEdge` instead of `UndirectedGraph.AddEdge` when adding edges to avoid adding negative-weight edges.