package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"bilgisayar", 500, "dell"})
	fmt.Println(product{name: "computer", unitPrice: 15000, brand: "casper"})
}

type product struct {
	name      string
	unitPrice float64
	brand     string
}
