package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) insertFront(node *node) {
	node.next = l.head
	l.head = node
	l.length++
}

func(l* linkedList) deleteValue(value int){
	if l.length == 0{
		fmt.Println("Linked List is Empty!")
	}
	if l.head.data == value{
		l.head = l.head.next
		l.length--
		return
	}
	las := l.head
	cur := l.head.next
	for cur != nil{
		if cur.data == value{
			las.next = cur.next
			l.length--
			return
		}
		las = cur
		cur = cur.next
	}
	fmt.Println("List doesnt not contains tha Value!")
}


func (l *linkedList) print(){
	if l.length == 0{
		fmt.Println("Linked List is Empty!")
	}
	cur := l.head
	for cur != nil{
		fmt.Print(cur.data , " ")
		cur = cur.next
	}
	println("")
}

func main() {
	list := linkedList{length: 0, head: nil}
	node1 := &node{data: 1}
	node2 := &node{data: 2}
	node3 := &node{data: 3}
	node4 := &node{data: 4}
	node5 := &node{data: 5}
	node6 := &node{data: 6}


	list.insertFront(node1)
	list.insertFront(node2)
	list.insertFront(node3)
	list.insertFront(node4)
	list.insertFront(node5)
	list.insertFront(node6)

	list.print()
	fmt.Println(list)
	list.deleteValue(100)
	list.print()
	fmt.Println(list)


}
