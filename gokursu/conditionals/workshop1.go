package conditionals

func Demo3() {

	var sayi1, sayi2, sayi3 int = 8, 2, 3

	var enBuyuk int = sayi1

	if sayi2 > enBuyuk {
		enBuyuk = sayi2

	}
	if sayi3 > enBuyuk {

		enBuyuk = sayi3

	}

	println("en büyük sayı", enBuyuk)
}
