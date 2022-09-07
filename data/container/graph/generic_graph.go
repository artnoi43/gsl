package graph

// GenericGraph represents what an gsl graph should look like.
// It is not intended to be used by production code, but more like an internal building block for gsl graphs.
// It's not minimal, and was designed around flexibility and coverage.
// This interface can be used for both unweighted and weighted graph (see wgraph package).
type GenericGraph[
	N any, // Type for graph node
	E any, // Type for graph edge
	M any, // Type representing the edge implementation of the graph, typically map[N]E.
	W any, // Type for graph edge weight or node values
] interface {
	// SetDirection sets the directionality of the graph.
	SetDirection(value bool)
	// HasDirection
	HasDirection() bool
	AddNode(node N)
	GetNodes() []N
	// AddEdge adds an edge to the graph. If the graph is directional, then AddEdge will only adds edge from n1 to n2.
	// If the graph is undirectional, then both n1->n2 and n2->n1 edges are added.
	AddEdge(n1, n2 N, weight W) error
	// GetNodeEdge takes in a node, and returns the connections from that node. If the graph is unweighted,
	// the so-called "edges" are just a list of nodes. If the graph is weighted, then the edges returned are
	// slice of the actual edges.
	GetNodeEdges(node N) []E
	// GetEdges returns all the edges in the graph.
	// In the actual implementations provided by gsl, the type M differs between unweighted and weighted graphs.
	// If the graph is weighted, then the returned type M is usually a map of node to its edges,
	// otherwise if the graph has unweighted edges, It returns a slice of nodes accessible by Node.
	GetEdges() M
}
