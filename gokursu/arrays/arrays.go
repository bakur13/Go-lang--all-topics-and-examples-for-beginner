package arrays

import "fmt"

func Demo1() {

	var cities [4]string

	cities[0] = "london"
	cities[1] = "Diyarbakır"
	cities[2] = "Berlin"
	cities[3] = "Dubai"

	for i := 0; i < 4; i++ {
		fmt.Println(cities[i])

	}

}
