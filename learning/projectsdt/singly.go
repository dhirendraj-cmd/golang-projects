package projectsdt

import "fmt"

type Node struct{
	data int
	next *Node
}

// constructor
func NewNode(data int) *Node {
	node := Node{
		data: data,
		next: nil,
	}
	return &node
}


// linkedlist struct
type SinglyLL struct{
	head *Node
}

// linkedlist constructor
func NewSinglyLL() *SinglyLL {
	return &SinglyLL{
		head: nil,
	}
}

func (list *SinglyLL) IsEmpty() bool{
	return list.head == nil
}

// Insert at End
func (list *SinglyLL) InsertAtEnd(newnode *Node) {
	if list.IsEmpty(){
		// fmt.Println("LIST IS EMPTY")
		list.head = newnode
		// fmt.Println("New node>> ", list.head.data)
	} else{ 
		lastNode := list.head
		for {
			if lastNode.next == nil{
				break
			} else {
				lastNode = lastNode.next
			}
		}
		lastNode.next = newnode
	}
}

// printing whole list
func (list *SinglyLL) PrintList(){

	if list.IsEmpty(){
		fmt.Println("List is Empty")
		return
	} else {
		currNode := list.head
		for currNode != nil {
			fmt.Print(currNode.data)
			if currNode.next != nil{
				fmt.Print("->")
			}
			currNode = currNode.next
		}
		fmt.Println()
	}

}

func LLops(){

	list := NewSinglyLL()
	
	node1 := NewNode(1)
	list.InsertAtEnd(node1)
	node2 := NewNode(2)
	list.InsertAtEnd(node2)
	node3 := NewNode(3)
	list.InsertAtEnd(node3)
	node4 := NewNode(4)
	list.InsertAtEnd(node4)
	node5 := NewNode(5)
	list.InsertAtEnd(node5)
	node6 := NewNode(6)
	list.InsertAtEnd(node6)

	list.PrintList()


}
