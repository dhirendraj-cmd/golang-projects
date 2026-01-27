package projectsdt

import "fmt"

type GenericNode[T any] struct{
	Data T
	Right *GenericNode[T]
	Left *GenericNode[T]
}

type BinaryTree[T any] struct{
	Root *GenericNode[T]
}

func NewGenericNode[T any](data T) *GenericNode[T]{
	return &GenericNode[T]{
		Data: data,
	}
}

func BuildGenericTree() *BinaryTree[int]{
	root := NewGenericNode(1)
	root.Left = NewGenericNode(2)
	root.Right = NewGenericNode(3)

	return &BinaryTree[int]{
		Root: root,
	}
}

func (g *GenericNode[T]) PreOrder(visit func(T)){
	if g==nil{
		return
	}

	visit(g.Data)
	g.Left.PreOrder(visit)
	g.Right.PreOrder(visit)
}
func (g *GenericNode[T]) PostOrder(visit func(T)){
	if g==nil{
		return
	}

	g.Left.PostOrder(visit)
	g.Right.PostOrder(visit)
	visit(g.Data)
}
func (g *GenericNode[T]) InOrder(visit func(T)){
	if g==nil{
		return
	}

	g.Left.InOrder(visit)
	visit(g.Data)
	g.Right.InOrder(visit)
}


func GenericCall(){
	tree := BuildGenericTree()


	fmt.Print("PreOrder Traversal: ")
	tree.Root.PreOrder(func(i int) {
		fmt.Printf("%d ", i)
	})
	fmt.Println()
	
	fmt.Print("PostOrder Traversal: ")
	tree.Root.PostOrder(func(i int) {
		fmt.Printf("%d ", i)
	})
	fmt.Println()

	fmt.Print("InOrder Traversal: ")
	tree.Root.InOrder(func(i int) {
		fmt.Printf("%d ", i)
	})
	fmt.Println()

}

