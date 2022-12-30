package conditionals

func Demo1() {

	var hesap int = 1000
	var cekme int = 100

	if cekme > hesap {
		println("Paran yok o kadar fakir herif")
	}

	if cekme <= hesap {
		println("Paranız hazırlanıyor")

		hesap = hesap - cekme
		println("Hesabınızda kalan toplam paranız", hesap)
	}

}
