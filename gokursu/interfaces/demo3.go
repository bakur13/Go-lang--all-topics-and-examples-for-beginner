package interfaces

import "fmt"

func dogrula(i interface{}) {

	sayi, ok := i.(int)
	fmt.Println(sayi, ok)

}
func Demo3() {
	var sayi1 interface{} = "nurşani"
	dogrula((sayi1))

	var sayi2 interface{} = 5
	fmt.Println(sayi2)
}
