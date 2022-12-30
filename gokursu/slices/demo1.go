package slices

import "fmt"

func Demo3() {

	cities := []string{"london", "paris", "silopi"}

	//fmt.Println(cities)
	citiescopy := make([]string, len(cities))
	copy(citiescopy, cities)
	//	fmt.Println(citiescopy)

	cities = append(cities, "berlin")
	//fmt.Println(cities, citiescopy)
	fmt.Println(cities[:2])

}
