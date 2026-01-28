package goroutines

import (
	"fmt"
	"time"
)

func RunTask(id int){
	fmt.Println("Running Task: ", id)
}

func CallTask(){
	for i:= range 10{
		go RunTask(i)
	}

	time.Sleep(time.Second*3)
}
