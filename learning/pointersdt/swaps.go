package pointersdt

import "fmt"

func Swap(x *int, y *int){
	var temp int

	temp=*x
	*x=*y
	*y=temp
}


func SwapVar(){
	var m,n int

	fmt.Print("Enter Numbers to Swap: ")
	fmt.Scanf("%d %d", &m, &n)

	fmt.Println("Before swap: ", m, n)

	Swap(&m, &n)

	fmt.Println("After Swap: ", m, n)

}
