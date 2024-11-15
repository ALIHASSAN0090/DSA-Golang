package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func main() {

	list := &LinkedList{}
	list.AddNode(10)
	list.AddNode(20)
	list.AddNode(30)

	list.PrintList()
}

func (list *LinkedList) PrintList() {
	current := list.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}

}

func (list *LinkedList) AddNode(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}
