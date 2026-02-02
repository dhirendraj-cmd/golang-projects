package channels

import (
	"fmt"
	"time"
)


func queueProcessing(emailChan chan string, done chan bool){
	defer func(){
		done <- true
		fmt.Println("Data Processed")
		}()

	for email := range emailChan{
		fmt.Println("sending email to ", email)
		time.Sleep(time.Second/2)
	}
	

}


func EmailSending(){
	emailChan := make(chan string, 100)
	done := make(chan bool)

	go queueProcessing(emailChan, done)

	for i:=range 100{
		emailChan <- fmt.Sprintf("%d@example.com", i+1)
	}

	fmt.Println("Added to Queue.... ")

	// closing channel
	close(emailChan)
	<-done

}

