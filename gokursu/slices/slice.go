package slices

import (
	"fmt"
)

func Demo1() {

	names := make([]string, 3)

	names[0] = "Hamza"
	names[1] = "Uslu"
	names[2] = "Silopi"
	names = append(names, "bakuri", "software engineer")
	fmt.Println(names)
	fmt.Println(len(names))
}
