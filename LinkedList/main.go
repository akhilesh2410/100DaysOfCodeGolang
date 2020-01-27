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
	myList.ShowList()
	myList.Delete(4)
	myList.ShowList()
}
