package slices

import "fmt"

func Demo2() {

	var ai [10]int
	var af [10]float64
	var as [10]string

	ai[5] = 500
	af[5] = 500.789
	as[5] = "Hi"

	for i := 0; i < len(ai); i++ {
		fmt.Printf("%d,", ai[i])
	}

	fmt.Println("\n----------------")
	for i := 0; i < len(af); i++ {
		fmt.Printf("%f,", af[i])
	}

	fmt.Println("\n----------------")
	for i := 0; i < len(as); i++ {
		fmt.Printf("%v,", as[i])
	}

	weekdays := [...]string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}
	for index := range weekdays {
		fmt.Println(weekdays[index])
	}
	fmt.Println(weekdays[0])
	fmt.Println("------------------EXAMPLE-------------------------------")

}

func Demo4() {

	var name1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	var name2 = [...]string{"Hamza", "Uslu", "software", "Devoloper", "Silopi", "ist Medeniyet universty", "26 years old", "abc", "123"}
	var name3 [10]float32

	for i := 0; i < len(name1); i++ {

		fmt.Println("  ", name1[i])

	}

	for i := 0; i < len(name2); i++ {

		fmt.Println("  ", name2[i])

	}
	for i := 0; i < len(name3); i++ {

		fmt.Println("  ", name3[i])

	}
	println("---------------EXAMPLE---------------")
}
