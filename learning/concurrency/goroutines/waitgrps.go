package goroutines

import (
	"fmt"
	"sync"
)


func task(id int, w *sync.WaitGroup){
	defer w.Done()
	fmt.Println("running task", id)
}


func CallTask1(){
	var wg sync.WaitGroup

	for i:= range 10{
		wg.Add(2)
		go task(i, &wg)
		go task(i+1, &wg)
	}

	wg.Wait()
}