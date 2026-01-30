package channels

import (
	"math/rand"
	"fmt"
	"sync"
	"time"
)



func GenerateRandomString(length int) string{

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randString := make([]byte, length)

	for i:= range length{
		randString[i] = charset[rand.Intn(len(charset))]
	}

	return string(randString)

}

func ProcessMsg(processChan chan string, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Printing Proces msg")
	fmt.Println(<-processChan)

}


func CallMsg(){

	var wg sync.WaitGroup

	// wg.Add(1)
	processChan := make(chan string, 1)
	// go ProcessMsg(processChan, &wg) // Start the Receiver First
	// processChan <- "Learning Golang channels!!"

	
	for range 100 {
		wg.Add(1)
		go ProcessMsg(processChan, &wg)
		processChan <- GenerateRandomString(5) + GenerateRandomString(8)
		time.Sleep(time.Second*1)
	}

	wg.Wait()

}
