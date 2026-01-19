package projectsdt

import "fmt"

type DNode struct{
	data int
	next *DNode
	prev *DNode
}

func DNewNode(data int) *DNode{
	return &DNode{
		data: data,
		next: nil,
		prev: nil,
	}
}

type Doubly struct{
	head *DNode
}

func NewDoubly() *Doubly{
	return &Doubly{
		head: nil,
	}
}

func (dll *Doubly) IsEmpty() bool{
	return dll.head == nil
}

func (dll *Doubly) InsertAtBegining(newnode *DNode){
	if dll.IsEmpty(){
		dll.head = newnode
		return
	} 
	temp := dll.head
	dll.head = newnode
	newnode.next = temp
	temp.prev = newnode

	// return newnode
	
}
func (dll *Doubly) InsertAtEnd(newnode *DNode){
	if dll.IsEmpty(){
		dll.head = newnode
		return 
	} 
	tempNode := dll.head
	for tempNode.next != nil{
		tempNode = tempNode.next
	}

	tempNode.next = newnode
	newnode.prev = tempNode
}

func (dll *Doubly) display(){
	if dll.IsEmpty(){
		fmt.Println("Empty Lsit!!!")
		return
	} else {
		currNode := dll.head
		for currNode != nil{
			fmt.Print(currNode.data)
			if currNode.next != nil{
				fmt.Print("<->")
			}
			currNode=currNode.next
		}
		fmt.Println()
	}
}

func DLLOps(){
	dll := NewDoubly()

	node1:=DNewNode(1)
	dll.InsertAtEnd(node1)
	dll.display()
	
	node2:=DNewNode(2)
	dll.InsertAtEnd(node2)
	dll.display()
	node3:=DNewNode(3)
	dll.InsertAtEnd(node3)
	dll.display()
	node4:=DNewNode(4)
	dll.InsertAtEnd(node4)
	dll.display()
	// node3:=DNewNode(3)

	node32:=DNewNode(32)
	dll.InsertAtBegining(node32)
	dll.display()
}