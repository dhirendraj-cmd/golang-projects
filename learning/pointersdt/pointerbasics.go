package pointersdt


import "fmt"


func BasicsPointer(){
	var ptr *int
	fmt.Println(ptr, &ptr)

	num := 55

	ptr = &num
	fmt.Println(ptr)
	fmt.Println(*ptr, &ptr)
	
	*ptr = 34
	fmt.Println(ptr)
	fmt.Println(*ptr, num, &num)



}