package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert
func (node *Node) Insert(v int) {
	if node == nil {
		return
	}
	if v < node.Key {
		if node.Left == nil {
			node.Left = &Node{Key: 0, Left: nil, Right: nil}
			node.Left.Key = v
			return
		} else {
			node.Left.Insert(v)
		}
	}
	if v > node.Key {
		if node.Right == nil {
			node.Right = &Node{Key: 0, Left: nil, Right: nil}
			node.Right.Key = v
			return
		} else {
			node.Right.Insert(v)
		}
	}
}

// Search
func (node *Node) Search(v int) bool {
	if node == nil {
		return false
	}
	if node.Key == v {
		return true
	}
	if v < node.Key {
		return node.Left.Search(v)
	}
	return node.Right.Search(v)
}

func main() {
	tree := &Node{Key: 100, Left: nil, Right: nil}
	tree.Insert(50)
	tree.Insert(75)
	tree.Insert(45)
	tree.Insert(21)
	tree.Insert(110)
	tree.Insert(150)
	tree.Insert(140)
	fmt.Println(tree)

	fmt.Println(tree.Search(75))
	fmt.Println(tree.Search(150))
	fmt.Println(tree.Search(21))
	fmt.Println(tree.Search(76))
}
