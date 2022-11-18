package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) Insert_end(val int) {
	//m:= node {data:val}
	m := node{}
	m.data = val
	if l.head == nil {
		l.head = &m
		l.length++
		return
	}

	temp := l.head
	for i := 0; i < l.length; i++ {
		if temp.next == nil {
			temp.next = &m
			l.length++
			return
		}
		temp = temp.next

	}
}

func (l *LinkedList) Insert_At(val int) {
	m := node{}
	m.data = val
	if l.length == 0 {
		l.head = &m
		l.length++
		return
	}
	temp := l.head
	m.next = temp
	l.head = &m
	l.length++
	return

}

func (l *LinkedList) Search(val int) {
	temp := l.head
	for i := 0; i < l.length; i++ {
		if temp.data == val {
			fmt.Println("paiyan pudichutan")
			return
		}
		temp = temp.next
	}
	fmt.Println("dhathiii")

}

func (l *LinkedList) Get_At(pos int) *node {
	temp := l.head
	for i := 0; i < pos; i++ {
		temp = temp.next
	}
	return temp
}
func (l *LinkedList) Insert_Any(pos int, val int) {
	m := node{}
	m.data = val
	if l.length == 0 {
		l.head = &m
		l.length++
		return
	}
	curr := l.Get_At(pos)
	m.next = curr
	prev := l.Get_At(pos - 1)
	prev.next = &m
	l.length++

}
func (l *LinkedList) Delete_any(pos int) {
	if l.length == 0 {
		return
	}
	prev := l.Get_At(pos - 1)
	next := l.Get_At(pos + 1)
	prev.next = next
	l.length--
}
func (l *LinkedList) Delete_First() {
	temp := l.head
	fmt.Println("Deleted element ", temp.data)

	l.head = temp.next
	l.length--

}

func (l *LinkedList) Delete_end() {
	temp := l.head
	for i := 0; i < l.length-1; i++ {
		temp = temp.next
	}
	temp.next = nil
	l.length--
	return

}

func (l *LinkedList) Display() {
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
	list := LinkedList{}
	list.Insert_end(10)
	list.Insert_end(20)
	list.Insert_end(30)
	list.Insert_end(40)
	list.Display()
	fmt.Println()
	list.Insert_Any(2, 25)
	list.Display()
	list.Delete_any(2)
	fmt.Println()
	list.Display()

}
