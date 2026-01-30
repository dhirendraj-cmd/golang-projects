package channels

import "fmt"

func DeadLockChan(){
	msgChan := make(chan string)

	msgChan <- "Learning GoLang" // blocking

	msgVal := <- msgChan

	fmt.Println(msgVal)
}

/*
	msgchan here is a channel sending string value to msgVal, but channels are blocking so deadlock condition arise which is solved in messages.go

	Above function is example of unbuffered channel
*/