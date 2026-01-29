package channels

import (
	"fmt"
	"sync"
	// "time"

)

func ProcessMsg(processChan chan string, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Printing Proces msg")
	fmt.Println(<-processChan)

}


func CallMsg(){

	var wg sync.WaitGroup

	wg.Add(1)
	processChan := make(chan string, 1)
	go ProcessMsg(processChan, &wg)
	processChan <- "Learning Golang channels!!"
	

	// time.Sleep(time.Second*1)
	wg.Wait()

}
