package arrays

import "fmt"

func Demo3() {

	var number [2][3]int

	number[0][0] = 1
	number[0][1] = 3
	number[0][2] = 5
	number[1][0] = 4
	number[1][1] = 6
	number[1][2] = 8

	for i := 0; i < 2; i++ {

		for j := 0; j < 3; j++ {

			fmt.Print(number[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}

}
