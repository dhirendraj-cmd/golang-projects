package pointersdt

import (
	"fmt"
)


func ArrPointers(){
	var n int
	var ptr *int
	fmt.Print("Enter Size of Array: ")
	fmt.Scanf("%d", &n)
	
	if n==0{
		return
	}

	arr := make([]int, n)

	for i:=0; i<n; i++{
		ptr = &arr[i]
		fmt.Printf("Enter element at position %d: ", i)
		fmt.Scanf("%d", ptr)
	}
	fmt.Println(arr)

	fmt.Println("\nPrinting using pointer")
	for i:=0;i<n;i++{
		ptr = &arr[i]
		fmt.Println(*ptr)
	}
}
