package main

import (
	"fmt"
	// "learning/learning/projectsdt/cli"
	// "learning/learning/interfacesdt"
	"learning/learning/projectsdt"
	"time"
)


func main(){
	start := time.Now()
	fmt.Println("Calling from Main!!!")
	// structsdt.StudentInfo()
	// structsdt.OrderInfo()
	// structsdt.AnonStructs()
	// structsdt.OdersInfo()

	// pointers
	// pointersdt.BasicsPointer()
	// a, b := 5, 8
	// fmt.Println(a,b)
	// pointersdt.AddNos(&a, &b)
	// pointersdt.SwapVar()
	// fmt.Println(a,b)

	// pointersdt.ArrPointers()
	// projectsdt.LLops()
	// projectsdt.DLLOps()

	// todos project
	// cli.CreateTask()

	// interfaces
	// interfacesdt.Bookshelf()

	// binary tree
	// projectsdt.TreeCall()
	projectsdt.GenericCall()


	end := time.Since(start)
	fmt.Println("End time is >>> ", end)
}

