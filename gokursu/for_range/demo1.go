package for_range

import "fmt"

func Demo1() {
	cities := []string{"London", "Paris", "Berlin", "Diyarbakır"}

	for i := 0; i < len(cities); i++ {
		println(cities[i])
	}

	for i, city := range cities {
		fmt.Print(i)
		fmt.Println(" ", city)
	}

	shr := map[string]string{"London": "England", "Paris": "France", "Berlin": "Germany"}
	for i, shrr := range shr {
		fmt.Print(i)
		fmt.Println(" ", shrr)

	}
}
