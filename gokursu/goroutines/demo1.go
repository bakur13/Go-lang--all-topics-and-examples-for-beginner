package goroutiness

import (
	"fmt"
	"time"
)

func Doublenumbers() {
	for i := 0; i < 10; i += 2 {
		fmt.Println(i)
		time.Sleep(1 * time.Millisecond)
	}

}
func Oddnumbers() {
	for i := 1; i < 10; i += 2 {
		fmt.Println(i)
		time.Sleep(1 * time.Millisecond)
	}

}
