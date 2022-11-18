package main

import "fmt"

type cnode struct {
	data int
	next *cnode
}
type CLinkedList struct {
	head   *cnode
	length int
}

func (l *CLinkedList) Insert_At(val int) {
	y := cnode{}
	y.data = val
	if l.head == nil {
		l.head = &y
		l.length++
		return
	}
	temp := l.head
	for i := 0; i < l.length; i++ {
		if temp.next == nil {
			temp.next = &y
			l.length++
			return
		}
		temp = temp.next
	}

}
func (l *CLinkedList) Display() {
	if l.length == 0 {
		fmt.Println("empty node")
	}
	temp := l.head
	var i int
	for i = 0; i < l.length; i++ {
		fmt.Printf("%d ->", temp.data)
		temp = temp.next
	}

}

func main() {

	list := CLinkedList{}
	list.Insert_At(10)
	list.Insert_At(20)
	list.Insert_At(30)
	list.Insert_At(40)
	list.Display()

}
