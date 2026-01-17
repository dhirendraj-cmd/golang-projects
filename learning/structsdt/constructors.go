package structsdt

import "fmt"

type Order struct{
	id int
	amount float64
	status string
}

// constructor
func NewOrder(id int, amount float64, status string) *Order{
	myOrder := Order{
		id: id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

func OrderInfo(){
	orderDetails := NewOrder(1, 342.65, "received")
	fmt.Println(orderDetails)
}