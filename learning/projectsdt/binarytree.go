package projectsdt

import "fmt"


type BasicNode struct{
	data int
	right *BasicNode
	left *BasicNode
}

func NewBasicNode(data int) *BasicNode{
	return &BasicNode{
		data: data,
		right: nil,
		left: nil,
	}
}

func (b *BasicNode) BuildBasicTree() {
	b.left = NewBasicNode(2)
	b.right = NewBasicNode(3)
	b.left.left = NewBasicNode(4)
	b.right.right = NewBasicNode(5)

}

func (b *BasicNode) PreOrder(){
	if b==nil{
		return
	}
	fmt.Printf("%d ", b.data)
	b.left.PreOrder()
	b.right.PreOrder()
}
func (b *BasicNode) PostOrder(){
	if b==nil{
		return
	}
	b.left.PostOrder()
	b.right.PostOrder()
	fmt.Printf("%d ", b.data)
}
func (b *BasicNode) InOrder(){
	if b==nil{
		return
	}
	b.left.InOrder()
	fmt.Printf("%d ", b.data)
	b.right.InOrder()
}

func TreeCall(){
	root := NewBasicNode(1)
	root.BuildBasicTree()

	fmt.Print("Inorder Traversal: ")
	root.InOrder()
	fmt.Println()
	
	fmt.Print("PreOrder Traversal: ")
	root.PreOrder()
	fmt.Println()

	fmt.Print("PostOrder Traversal: ")
	root.PostOrder()
	fmt.Println()

}

