package string_fonk

import (
	"fmt"
	"strings"
)

func Demo1() {
	s := []string{"foo", "bar", "baz"}
	isim := "Hamza uslu"

	fmt.Println(strings.Contains(isim, "n"))
	fmt.Println(strings.Count(isim, "a"))
	fmt.Println(strings.ToLower(isim))
	fmt.Println(strings.ToUpper(isim))
	fmt.Println(strings.Index(isim, "a"))
	fmt.Println(strings.Join(s, "-"))
	fmt.Println(strings.Repeat(isim, 2))
	fmt.Println(strings.Split(isim, "a"))
	fmt.Println(strings.Replace(isim, "a", "A", -1))
}
