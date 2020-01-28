package main

import dll "100DaysOfCodeGolang/LinkedList/DoublyLinkList/doublylist"

func main(){
	mylist := &dll.DList{}
	mylist.Insert(1)
	mylist.Insert(2)
	mylist.Insert(3)
	mylist.Insert(4)
	mylist.Insert(5)
	mylist.InsertAtPos(15,3)
	mylist.ShowList()
	mylist.Delete(4)
	mylist.ShowList()
}
