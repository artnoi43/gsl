package graph

import "github.com/pkg/errors"

var ErrEdgeWeightNotNull = errors.New("found edge weight in unweighted graph")
