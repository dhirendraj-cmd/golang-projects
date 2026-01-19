package pointersdt

import "fmt"

func AddNos(a *int, b *int)  {
	p1, p2 := a, b

	sum := *p1+*p2

	fmt.Println(sum)

	*p1 = 23
	*p2 = 32

	fmt.Println(*p1+*p2)

	// user input values in pointers
	fmt.Print("Enter numbers to add: ")
	fmt.Scanf("%d%d", p1, p2)
	p3, p4 := p1, p2

	newSum := *p3+*p4

	fmt.Println(*p3, *p4, newSum)
}

