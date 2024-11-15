package main

import (
	"fmt"
)

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

	fmt.Printf("Initial list:\n")
	list.PrintList()
	list.AddNodeToFront(30)
	fmt.Printf("List after adding 30 to the front:\n")
	list.PrintList()
	list.AddNode(70)
	fmt.Printf("List after adding 70 to the front:\n")
	list.PrintList()
	list.AddtoIndex(3, 90)
	fmt.Printf("List after adding 90 at index 3:\n")
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

func (list *LinkedList) AddNodeToFront(data int) {
	newNode := &Node{data: data}
	newNode.next = list.head
	list.head = newNode
}

func (list *LinkedList) AddtoIndex(index, data int) {

	if index < 0 {
		fmt.Print("invalid index")
	}

	if index == 0 {
		list.AddNodeToFront(data)
	}

	newNode := &Node{data: data}
	current := list.head

	for i := 0; i < index-1 && current != nil; i++ {
		current = current.next
	}

	if current == nil {
		fmt.Print("index out of bounds")
		return
	}
	newNode.next = current.next
	current.next = newNode

}
