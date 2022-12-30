package for_range

func Demo2() {
	total := 0
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, number := range numbers {
		if number%3 != 0 {
			total = total + number

		}
	}
	println("total value", total)
}

// 3 sayısına tam bölünmeyen sayıların toplamınu bulan fonksiyon
