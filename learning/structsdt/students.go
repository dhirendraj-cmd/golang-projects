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
	vansh := StudentDetails{1, "Vansh", "vansh@gmail.com", "Mech", 330, 8.4, "distinction"}
	fmt.Println("Vansh details before: ", vansh)
	vansh.result = "First"
	fmt.Println("Vansh details after: ", vansh)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	
	vidyut := StudentDetails{2, "Vidyut", "vids@exp.com", "CS", 339, 8.5, "distinction"}
	fmt.Println("Vidyut details are: ", vidyut)

	dj := StudentDetails{2, "Dhiren", "dj@exp.com", "IT", 223, 4.5, "fail"}
	fmt.Println("DJ details are: ", dj)

}

