package channels

import (
	"fmt"
)

func Doublenumbers(doublesnumcn chan int) {
	total := 0
	for i := 0; i < 10; i += 2 {
		fmt.Println(i)
		total = total + i

	}
	fmt.Println("Doubles total", total)
	doublesnumcn <- total
}
func Oddnumbers(oddnumscn chan int) {
	total := 0
	for i := 1; i < 10; i += 2 {
		fmt.Println(i)
		total = total + i

	}
	fmt.Println("Odds total", total)
	oddnumscn <- total
}
