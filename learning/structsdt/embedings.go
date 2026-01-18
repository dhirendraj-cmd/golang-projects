// inheritance of structs basically composition it follows stating it follows has-a relationship

package structsdt

import (
	"fmt"
	"time"
)

type Customer struct{
	name string
	phone string
}


type Orders struct{
	id int
	amount float64
	status string
	createdAt time.Time
	Customer
}

func OdersInfo(){
	newOrder := Orders{
		id: 1,
		amount: 432.55,
		status: "Paid",
		Customer: Customer{
			name: "Holmes",
			phone: "234123567",
		},
	}

	fmt.Println(newOrder)
}
