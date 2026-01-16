package structsdt

import "fmt"

type StudentDetails struct{
	rollNo int
	name string
	email string
	department string
	marks int
	cgpa float64
	result string
}

func (sd *StudentDetails) ChangeResult(result string)  {
	sd.result = result
}

func StudentInfo(){
	students := []StudentDetails{}

}

func AddStudents()

