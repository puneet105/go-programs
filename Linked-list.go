package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (x *LinkedList) Insert(value int) {

	newNode := &Node{data: value, next: nil}
	if x.head == nil {
		x.head = newNode
		return
	}
	current := x.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) Delete(value int) {
	if ll.head == nil {
		fmt.Println("List is empty")
		return
	}

	if ll.head.data == value {
		ll.head = ll.head.next
		return
	}

	current := ll.head
	var prev *Node

	for current != nil && current.data != value {
		prev = current
		current = current.next
	}
	if current != nil {
		prev.next = current.next
	} else {
		fmt.Println("Value not found in the list")
	}
}

func main() {
	x := LinkedList{}
	x.Insert(10)
	x.Insert(20)
	x.Insert(30)
	x.Insert(40)

	fmt.Println("After Insertion")
	current := x.head
	for current != nil {
		fmt.Printf("%d\t%v\n", current.data,&current.next)
		current = current.next
	}
	x.Delete(40)
	fmt.Println("After Deletion")
	current = x.head
	for current != nil {
		fmt.Printf("%d\t%v\n", current.data,&current.next)
		current = current.next
	}

}
