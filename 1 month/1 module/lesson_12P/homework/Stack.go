package main

import "fmt"

type Stack struct {
	items []int
}

// Push
func (s *Stack) push(item int){
	s.items = append(s.items,item)
}

// Pop
func (s *Stack) pop()int{
	length:=len(s.items)-1
	removeItem := s.items[length]
	s.items = s.items[:length]
	return removeItem
}

func main(){
	var stack Stack
	stack.push(5)
	stack.push(12)
	stack.push(4)
	stack.push(23)
	stack.push(124)
	stack.push(14)
	stack.push(45)
	stack.push(6)
	fmt.Println(stack)
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack)
}