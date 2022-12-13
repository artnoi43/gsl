package wgraph

import "github.com/pkg/errors"

var (
	ErrConnExists                 = errors.New("node connection already exists")
	ErrDijkstraNegativeWeightEdge = errors.New("a Dijkstra edge's weight must not be negative")
)

func wrapErrConnExists[W Weight](n1, n2 NodeWeighted[W]) error {
	return errors.Wrapf(ErrConnExists, "existing edge between node %s to %s", n1.GetKey(), n2.GetKey())
}
