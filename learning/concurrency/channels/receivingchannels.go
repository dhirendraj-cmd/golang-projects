package channels

import "fmt"


// receive
func sum(res chan int, a int, b int){
	ans := a+b
	res <- ans
}

func Add(){
	res := make(chan int)

	go sum(res, 4, 5)
	result := <- res // blocking

	fmt.Println(result)
}
