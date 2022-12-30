package error_handling

import (
	"fmt"
	"os"
)

// type assertion
func Demo1() {
	f, err := os.Open("demo1.txt")
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadı", pErr.Path)
			return
		}

	} else {
		fmt.Println(f.Name(), "Dosya açılıyor")
		return
	}

}
