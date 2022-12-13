# Package `graph`

This package provides basic building blocks for working with graphs.

There's the generic interface `Graph[N, E, W]`, which all other graph implementations must implement.

The algorithm code then operates on `Graph`. You can easily implement `Graph`
with your own types, as shown in [this Dijkstra example](./assets/example_flight_path/main.go).

Here are some of `Graph` implementations in this package and sub-packages:

1. `HashMapGraphV1[T]` implements `Graph[Node[T], Node[T], any]`
   This interface represents a simple, unweighted graph that lacks a real edge type, i.e. it uses `Node[T]` as its edge type.
   It is designed this way to solve some unweighted graph problems where edges are not usually needed, like BFS search.
   `HashMapGraphV1Impl` is the type implementing this interface.

2. `SafeGraph[N, E, W]` implements `Graph[N, E, W]`
   This type is only a mutex-enabled wrapper for `Graph[N, E, W]`. Most of the default constructor functions wraps their
   undelying graphs with `SafeGraph`, unless the constructor ends with `-Unsafe`.

3. `GraphWeighted[N NodeWeighted[W], E EdgeWeighted[W, NodeWeighted[W]], W]` implements `Graph[N, E, W]`
   This interface represents a weighted graph. It is subject to a lot of changes as of now.

4. `DijkstraGraph[T]` implements `GraphWeighted[NodeDijkstra[T], EdgeWeighted[T, NodeDijkstra[T]], T]`
   This type wraps any `GraphWeighted` and provides methods for finding Dijkstra's shortest path.
