package LinkedList

import (
	"fmt"
)

type LinkList struct {
	data int
	next *LinkList
}

func (node *LinkList) Insert(data int) {
	newNode := new(LinkList)
	newNode.data = data
	newNode.next = nil
	if node == nil {
		node = newNode
	} else {
		temp := node
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
	}
}
func (list *LinkList) ShowList() {
	fmt.Printf("head")
	for ; list != nil; list = list.next {
		fmt.Printf("---->")
		fmt.Printf("%d", list.data)
	}
	fmt.Println("---->nil")
}
func (list *LinkList) Delete(data int) {
	if list == nil {
		fmt.Println("List is empty")
		return
	}
	temp := list
	for list.data != data {
		temp = list
		list = list.next
	}
	temp.next = list.next
}
