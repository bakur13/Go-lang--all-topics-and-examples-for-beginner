package structs

import "fmt"

type customer struct {
	name     string
	lastname string
	age      int
}

func (c customer) save() {
	fmt.Println("Eklendi", c.name)
}
func Demo2() {
	c := customer{name: "Hamza", lastname: "uslu", age: 26}
	c.save()
}
