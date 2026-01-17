package structsdt

import "fmt"

func AnonStructs(){
	dept := struct{
		deptId int
		deptName string
	} {
		deptId: 22,
		deptName: "CS",
	}

	fmt.Println(dept)
}