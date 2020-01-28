package main

import (
	ll "100DaysOfCodeGolang/LinkedList/list"
)

func main() {
	myList := &ll.LinkList{}
	myList.Insert(1)
	myList.Insert(2)
	myList.Insert(3)
	myList.Insert(4)
	myList.Insert(5)
	myList.InsertAtPos(15,3)
	myList.ShowList()
	myList.Delete(5)
	myList.ShowList()
	head,_ := myList.Reverse()
	head.ShowList()

}
