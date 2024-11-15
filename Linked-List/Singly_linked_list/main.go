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
	list.DeleteFromStart()
	fmt.Printf("List after deleting from the start:\n")
	list.PrintList()
	list.DeleteFromEnd()
	fmt.Printf("List after deleting from the end:\n")
	list.PrintList()
	fmt.Printf("Is the list empty? %v\n", list.IsEmpty())
	fmt.Printf("Length of the list: %d\n", list.GetLength())
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

func (list *LinkedList) DeleteFromStart() {
	if list.head == nil {
		fmt.Println("List is empty, nothing to delete.")
		return
	}
	list.head = list.head.next //unlinked the node from the linked list
}

func (list *LinkedList) DeleteFromEnd() {
	if list.head == nil {
		fmt.Println("List is empty, nothing to delete.")
		return
	}
	if list.head.next == nil {
		list.head = nil
	}

	current := list.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
}

func (list *LinkedList) IsEmpty() bool {
	if list.head == nil {
		return true
	}
	return false
}

func (list *LinkedList) GetLength() int {
	if list.IsEmpty() {
		return 0
	}

	values := 0

	current := list.head
	for current != nil {
		current = current.next
		values += 1

	}

	return values

}
