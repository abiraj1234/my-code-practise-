package main

import "fmt"

type d_node struct {
	data int
	next *d_node
}

type stackLinkedList struct {
	head   *d_node
	length int
}

func (s *stackLinkedList) push(val int) {
	mn := d_node{}
	mn.data = val
	if s.head == nil {
		s.head = &mn
		s.length++
		return
	}

	temp := s.head
	for i := 0; i < s.length; i++ {
		if temp.next == nil {
			temp.next = &mn
			s.length++
			return

		}
		temp = temp.next
	}
}

func (s *stackLinkedList) pop() {
	temp := s.head
	for i := 0; i < s.length-1; i++ {
		temp = temp.next
	}
	temp.next = nil
	s.length--
	return
}

func (s *stackLinkedList) Display() {
	if s.length == 0 {
		fmt.Println("empty node")
	}
	temp := s.head
	for i := 0; i < s.length; i++ {
		fmt.Printf("%d ->", temp.data)
		temp = temp.next
	}
}

func main() {
	list := stackLinkedList{}

	list.push(10)
	list.push(20)
	list.push(30)
	list.push(40)
	list.push(50)
	list.Display()
	fmt.Println()
	list.pop()
	list.Display()

}
