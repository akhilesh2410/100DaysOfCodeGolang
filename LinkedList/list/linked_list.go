package list

import (
	"fmt"
)

type LinkList struct {
	data int
	next *LinkList
}
//Insert element at last of list
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
//Insert element at a specific position in list
func (node *LinkList) InsertAtPos(data,pos int) {
	if node == nil {
		fmt.Println("list is  empty")
		return
	}
	newNode := new(LinkList)
	newNode.data = data
	newNode.next = nil
	count := 1
	temp := node
	for (temp.next != nil && count < pos) {
		temp = temp.next
		count++
	}
	if count < pos {
		fmt.Println("Not enough element present in the list")
		return
	}
	newNode.next = temp.next
	temp.next = newNode
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

func (list *LinkList) Reverse()(head *LinkList, cur *LinkList) {

	if list == nil {
		fmt.Println("list is empty")
		return nil, nil
	}
	//var head *LinkList
	if list.next == nil {
		return list, list
	}
	head, temp := (list.next).Reverse()
	temp.next = list
	list.next = nil
	return head ,list
}
