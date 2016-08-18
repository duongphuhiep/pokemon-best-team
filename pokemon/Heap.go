package pokemon

// An NodeHeap is a min-heap of Nodes.
type NodeHeap []*Node

func (h *NodeHeap) Push(x *Node) {
	// Push and Pop use poNodeer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x)
}

func (h *NodeHeap) Pop() *Node {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
