package main

type binaryHeapNode[T comparable] struct {
	value T
}

func (n *binaryHeapNode[T]) GetValue() T {
	return n.value
}

type binaryHeap[T comparable] []binaryHeapNode[T]

func (h *binaryHeap[T]) AddNode(node binaryHeapNode[T]) {
	*h = append(*h, node)
}

func (h *binaryHeap[T]) GetNodes() []binaryHeapNode[T] {
	return *h
}

// Get
func (h *binaryHeap[T]) GetNodeEdges(parent binaryHeapNode[T]) []binaryHeapNode[T] {
	nodeEdges := make([]binaryHeapNode[T], 2)
	slice := *h
	for i, item := range slice {
		if item.GetValue() == parent.GetValue() {
			// item is parent node
			for j := 0; j < 2; j++ {
				nodeEdges[j] = slice[(2*i)+j+1]
			}
		}
	}

	return nodeEdges
}

func (h *binaryHeap[T]) AddEdge(n1, n2 binaryHeapNode[T], w T) error {
	return nil
}
