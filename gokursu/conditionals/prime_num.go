package conditionals

import (
	"fmt"
)

// function to find prime number
func Demo5() {

	var num int
	fmt.Println("Please  a number")
	fmt.Scanln(&num)
	moment := true
	for i := 2; i < num; i++ {

		if num%i == 0 {

			moment = false

		}
	}
	if moment == true {

		fmt.Println("Your number is a prime number")
	} else {
		fmt.Println("Your number is not prime")
	}
}
