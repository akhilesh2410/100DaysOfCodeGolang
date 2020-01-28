package doublylist

import "fmt"

type DList struct {
	data int
	prev *DList
	next *DList
}

func (node *DList) Insert(data int) {
	newNode := new(DList)
	newNode.data = data
	newNode.prev = nil
	newNode.next = nil
	if node == nil {
		node = newNode
	} else {
		temp := node
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
		newNode.prev = temp
	}
}

//Insert element at a specific position in list
func (node *DList) InsertAtPos(data,pos int) {
	if node == nil {
		fmt.Println("list is  empty")
		return
	}
	newNode := new(DList)
	newNode.data = data
	newNode.prev = nil
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
	temp.next.prev = newNode
	newNode.prev = temp
	temp.next = newNode
}
func (list *DList) ShowList() {
	fmt.Printf("head")
	for ; list != nil; list = list.next {
		fmt.Printf("<---->")
		fmt.Printf("%d", list.data)
	}
	fmt.Println("<---->nil")
}
func (list *DList) Delete(data int) {
	if list == nil {
		fmt.Println("List is empty")
		return
	}
	temp := list
	for list != nil && list.data != data {
		temp = list
		list = list.next
	}
	if list == nil {
		fmt.Println("Node not found with given data")
		return
	}
	temp.next = list.next
	list.next.prev = temp
}