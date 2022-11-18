package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}
type DoubleLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (l *DoubleLinkedList) Insert_end(val int) {
	x := Node{}
	x.data = val
	if l.length == 0 {
		l.head = &x
		l.tail = &x
		l.length++
		return
	}
	temp := l.tail
	x.prev = temp
	temp.next = &x
	l.tail = &x
	l.length++
}
func (l *DoubleLinkedList) Display() {
	if l.length == 0 {
		fmt.Println("ulla onum ila ")
		return
	}
	temp := l.head
	for i := 0; i < l.length; i++ {
		fmt.Printf("%d ", temp.data)
		temp = temp.next
	}
}
func (l *DoubleLinkedList) reverse() {
	if l.length == 0 {
		fmt.Println("ulla onum ila ")
		return
	}
	temp := l.tail
	for i := 0; i < l.length; i++ {
		fmt.Printf("%d ", temp.data)
		temp = temp.prev
	}
}
func (l *DoubleLinkedList) Delete_First() {
	if l.length == 0 {
		fmt.Println("hy ena kururkurunu pakkura ")
		return
	}
	temp := l.head
	fmt.Println("Deleted element is ", temp.data)
	l.head = temp.next
	temp.prev = nil
	l.length--

}
func (l *DoubleLinkedList) Delete_last() {
	if l.length == 0 {
		fmt.Println("hy ena kururkurunu pakkura ")
		return
	}
	temp := l.tail
	l.tail = temp.prev
	l.tail.next = nil
	l.length--
}
func (l *DoubleLinkedList) Search(val int) {
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

func (l *DoubleLinkedList) Get_At(pos int) *Node {
	if pos < 0 {
		return l.head
	}
	temp := l.head
	for i := 0; i < pos; i++ {
		temp = temp.next
	}
	return temp
}

func (l *DoubleLinkedList) Insert_Any(pos int, val int) {
	x := Node{}
	x.data = val
	if l.length == 0 {
		l.head = &x
		l.tail = &x
		l.length++
		return
	}
	curr := l.Get_At(pos)
	prev := l.Get_At(pos - 1)
	x.next = curr
	x.prev = prev
	prev.next = &x
	curr.prev = &x
	l.length++
}

func (l *DoubleLinkedList) Delete_any(pos int) {
	if l.length == 0 {
		return
	}
	prev := l.Get_At(pos - 1)
	next := l.Get_At(pos + 1)

	prev.next = next
	next.prev = prev
	l.length--
}

func (l *DoubleLinkedList) insert_first(val int) {
	x := Node{}
	x.data = val
	if l.length == 0 {
		l.head = &x
		l.tail = &x
		l.length++
		return
	}
	temp := l.head
	temp.prev = &x
	x.next = temp
	l.head = &x
	l.length++
}

func main() {
	list := DoubleLinkedList{}
	list.Insert_end(10)
	list.Insert_end(20)
	list.Insert_end(30)
	list.Insert_end(40)
	list.Insert_end(50)
	list.insert_first(-10)
	list.Display()
	fmt.Println()
	list.Insert_Any(2, 15)
	list.Display()
	fmt.Println()
	list.Delete_any(4)
	list.Display()
	list.Search(000)
	list.Display()
}
