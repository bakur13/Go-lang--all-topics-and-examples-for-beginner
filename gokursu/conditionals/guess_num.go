package conditionals

import "fmt"

func Demo4() {

	my := 0
	your := 0
	guess := 0

	fmt.Println("Please enter a number")
	fmt.Scanln(&my)
	fmt.Println("Guess the number I'm holding")
	fmt.Scanln(&your)
	//fmt.Print("bu sayi", my)

	for my != your {

		if my > your {
			println("Please enter a larger number")
			fmt.Scanln(&your)
			guess = guess + 1
		}
		if my < your {

			fmt.Println(" Please enter a smaller  number")
			fmt.Scanln(&your)
			guess = guess + 1
		}

	}
	if guess < 3 {
		fmt.Println("super")

	}

	guessnum := ""
	if guess > 0 && guess <= 3 {
		guessnum = "super"
	} else if guess <= 6 {

		guessnum = "GOOD :))"

	} else {

		guessnum = "EHHH NORMAL"
	}
	fmt.Printf("you found %v in the try You are   %v", guess, guessnum)

}
