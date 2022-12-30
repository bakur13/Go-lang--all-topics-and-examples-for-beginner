package goroutiness

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printX() {
	for i := 0; i < 10; i++ {
		fmt.Println(" *       *")
		fmt.Println(" **    *** ")
		fmt.Println(" ***   *** ")
		fmt.Println("  ******** ")
		fmt.Println("   *****   ")
		fmt.Println("    ***    ")
		fmt.Println("     *     ")
	}

}

func printy() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hamza")
	}
	wg.Done()
}

func Demo2() {

	wg.Add(1)
	go printy()

	wg.Wait()
	fmt.Println()
	printX()

}
