package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Print() {
	fmt.Println("Linked List: ")
	poiter := l.head
	for i := 0; i < l.length; i++ {
		fmt.Printf("%d ", poiter.data)
		poiter = poiter.next
	}
	fmt.Println()
}

// 1. AddFirst Method

func (l *LinkedList) AddFirst(n *Node) {

	if l.length == 0 {
		l.head = n
		l.length++

		return
	}

	old_list := l.head

	l.head = n
	l.length++
	l.head.next = old_list
	return
}

// 2. Remove First Method,
func (l *LinkedList) RemoveFirst() {
	l.head = l.head.next
	l.length--
}

// 3. GetAt Method
func (l *LinkedList) GetAt(position int) *Node {
	if position > l.length {
		return &Node{}
	}
	node := l.head
	for i := 1; i < position; i++ {

		node = node.next
		fmt.Println("i=", i)
	}

	return node

}

// Middle Element method
func (l *LinkedList) Middle_Element() *Node {

	node := l.head
	middle_element := 1
	if l.length%2 == 0 {
		middle_element = l.length / 2
	} else {
		middle_element = l.length/2 + 1
	}

	for i := 1; i < middle_element; i++ {
		node = node.next
		fmt.Println("i=", i, middle_element)
	}

	return node

}

/*********************************HOMEWORK*********************************/

// 2. Remove Middle Element
func (l *LinkedList) Remove_Middle_Element() {
	list_length := 0
	old_list := l.head
	// middle_element := l.head
	if l.length%2 == 0 {
		list_length = l.length / 2
		
	} else {
		list_length = l.length/2 + 1
	}

	for i := 1; i <= l.length; i++ {
	
		if list_length == i&&old_list.next!=nil{
			old_list = old_list.next 
			l.head.next = old_list
			fmt.Println("old",old_list)
		}else if list_length == i+1&&old_list.next!=nil{
				fmt.Println(i,"    ",list_length,"     ",old_list)
				old_list=old_list.next
				l.head.next=old_list
			}
	}
	l.head=old_list
	l.length--
	// l.head.next=middle_element.next
	l.Print()
	// l.length--

}

/*********************************HOMEWORK*********************************/

func main() {
	var linkedList = &LinkedList{head: nil, length: 0}
	linkedList.AddFirst(&Node{data: 2})
	linkedList.AddFirst(&Node{data: 4})
	linkedList.AddFirst(&Node{data: 5})
	linkedList.AddFirst(&Node{data: 53})
	linkedList.AddFirst(&Node{data: 54})
	linkedList.AddFirst(&Node{data: 55})
	linkedList.AddFirst(&Node{data: 56})
	linkedList.Print()
	// linkedList.RemoveFirst()
	// linkedList.RemoveFirst()
	// linkedList.RemoveFirst()
	linkedList.Remove_Middle_Element()
	linkedList.Print()
	// fmt.Println(linkedList.GetAt(1))
	// fmt.Println(linkedList.Middle_Element())

}
