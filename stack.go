package main

import "fmt"

type stack struct {
	array []int
}

func (s *stack) push(data int) {
	s.array = append(s.array, data)
}
func (s *stack) size() int {
	return len(s.array)
}

func (s *stack) print() {
	if s.size() == 0 {
		return
	}
	for _, i := range s.array {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func (s *stack) pop() {
	if s.size() == 0 {
		return
	}
	fmt.Println("sirapana alu nee", s.array[s.size()-1])
	s.array = s.array[0 : s.size()-1]

}

func main() {
	nonut := stack{}
	nonut.push(100)
	nonut.push(200)
	nonut.push(300)
	nonut.push(400)
	nonut.print()
	nonut.pop()
	nonut.print()
}
