package conditionals

import "fmt"

func Demo6() {

	var num1, num2, top1, top2 int = 0, 0, 0, 0

	println("please enter a number ")
	fmt.Scanln(&num1)
	println("Please enter second a number ")
	fmt.Scanln(&num2)

	for i := 1; i < num1; i++ {

		if num1%i == 0 {

			top1 = top1 + i
		}
	}

	for j := 1; j < num2; j++ {
		if num2%j == 0 {
			top2 = top2 + j
		}
	}

	if top1 == num2 && top2 == num1 {
		fmt.Println("these are friend numbers")
	} else {

		fmt.Printf("these %v %v aren't friend numbers", num1, num2)
	}

}
