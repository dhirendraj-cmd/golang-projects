package channels

import "fmt"

func BufferedChan(){
	// buffered channnels simply avoid blocking bcoz of using the capacity, so till that size it remains unblocked

	exmChan := make(chan string, 3)

	exmChan <- "this is first"
	exmChan <- "this is 2nd"
	exmChan <- "this is 3rd"


	fmt.Println(<-exmChan)
	fmt.Println(<-exmChan)
	fmt.Println(<-exmChan)


}
