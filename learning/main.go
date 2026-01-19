package main

import (
	"fmt"
	"learning/learning/pointersdt"
)


func main(){
	fmt.Println("Calling from Main!!!")
	// structsdt.StudentInfo()
	// structsdt.OrderInfo()
	// structsdt.AnonStructs()
	// structsdt.OdersInfo()

	// pointers
	// pointersdt.BasicsPointer()
	a, b := 5, 8
	fmt.Println(a,b)
	// pointersdt.AddNos(&a, &b)
	// pointersdt.SwapVar()
	// fmt.Println(a,b)

	pointersdt.ArrPointers()
}

