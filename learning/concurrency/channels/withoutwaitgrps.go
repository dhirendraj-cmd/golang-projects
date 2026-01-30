package channels

import "fmt"


// making go rountines wait without using sync.WaitGroup
func task(done chan bool){

	// below line will run at the end due to defer keyword
	defer func ()  {
		done <- true
	} ()

	fmt.Println("Processing data......", done)
}

func WaitWithoutWaitGroups(){
	done := make(chan bool)

	// start receiver first
	go task(done)

	// in below line blocking will happen since done is receiveing data from task function via channel
	<-done //blocking
}
