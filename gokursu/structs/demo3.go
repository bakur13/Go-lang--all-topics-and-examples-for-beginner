package structs

import "fmt"

func Demo3() {
	type employee struct {
		name      string
		age       int
		isMarried bool
	}

	var e1 employee
	e1.name = "Hamza"
	e1.age = 26
	e1.isMarried = false

	var e2 employee
	e2.name = "Kadir"
	e2.age = 67
	e2.isMarried = true

	e3 := employee{
		name:      "Hatice",
		age:       57,
		isMarried: true,
	}
	fmt.Println(e1)
	fmt.Printf("%#v\n", e1)
	fmt.Println(e2)
	fmt.Printf("%#v\n", e2)
	fmt.Println(e3)
	fmt.Printf("%#v\n", e3)
}
