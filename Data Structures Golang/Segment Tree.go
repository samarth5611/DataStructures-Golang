package main

import "fmt"

// struct for segment node
type Node struct {
	Sum   int
	Left  *Node
	Right *Node
}

// Build
func Build(a *[]int, l int, r int) *Node {
	if l == r {
		return &Node{Sum: (*a)[l], Left: nil, Right: nil}
	}
	mid := (l + r) / 2
	LeftNode := Build(a, l, mid)
	RightNode := Build(a, mid+1, r)
	return &Node{Sum: LeftNode.Sum + RightNode.Sum, Left: LeftNode, Right: RightNode}
}

// Query
func (node *Node) Query(queryLeftRange int, queryRightRange int, nodeLeftRange int, nodeRightRange int) int {
	if queryLeftRange <= nodeLeftRange && nodeRightRange <= queryRightRange {
		return node.Sum
	}
	if queryRightRange < nodeLeftRange || nodeRightRange < queryLeftRange {
		return 0
	}
	mid := (nodeLeftRange + nodeRightRange) / 2
	return node.Left.Query(queryLeftRange, queryRightRange, nodeLeftRange, mid) + node.Right.Query(queryLeftRange, queryRightRange, mid+1, nodeRightRange)
}

// Update
func (node *Node) Update(nodeLeftRange int, nodeRightRange int, index int, value int) {
	if nodeLeftRange == index && nodeRightRange == index {
		node.Sum = value
		return
	}
	if nodeLeftRange <= index && nodeRightRange >= index {
		mid := (nodeLeftRange + nodeRightRange) / 2
		node.Left.Update(nodeLeftRange, mid, index, value)
		node.Right.Update(mid+1, nodeRightRange, index, value)
		node.Sum = node.Left.Sum + node.Right.Sum
	}
}
func main() {
	a := []int{2, 3, 4, 5, 1, 2}
	N := len(a) - 1
	T := Build(&a, 0, N)
	fmt.Println(T)
	fmt.Println(T.Query(0, 2, 0, N))
	fmt.Println(T.Query(2, 4, 0, N))
	fmt.Println(T.Query(5, 5, 0, N))
	T.Update(0, N, 5, 5)
	fmt.Println(T.Query(5, 5, 0, N))

}
