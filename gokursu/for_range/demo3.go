package for_range

func Demo3() {

	books := map[string]string{"Hobbit": "JRR Tolkien", "Poetika": "İbn'i sina"}

	for k, v := range books {
		print(k)
		println(" ", v)
	}

}
