package main

import (
	"fmt"
)

type Node struct{
	data interface{}
	next *Node
}

type List struct{
	head *Node
}

func (l *List) Insert(data interface{}) {
	list := &Node{
		next: l.head,
		data: data,
	}
	l.head = list
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.data)
		list = list.next
	}
	fmt.Println()
}

func main() {
	link := List{}
	link.Insert(5)
	link.Insert(9)
	link.Insert(13)
	link.Insert(22)
	link.Insert(28)
	link.Insert(36)

	link.Display()

	fmt.Printf("%d", getNFromLast(link, 3))
}

func getNFromLast(list List, k int) (data interface{}){
	front := list.head

	for i:= 0; i < k; i++ {
		front = front.next
	}

	back := list.head

	for front != nil {
		front = front.next 
		back = back.next
	}

	return back.data
}