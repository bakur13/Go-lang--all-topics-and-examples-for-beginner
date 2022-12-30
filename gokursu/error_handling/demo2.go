package error_handling

import (
	"errors"
	"fmt"
)

func guess(gues int) (string, error) {
	ANumber := 79
	if gues < 1 || gues > 100 {
		return "", errors.New("lütfen 1 ila 100 arasında bir sayı giriniz")

	}

	if gues < ANumber {
		return "Lütfen daha büyük bir sayı giriniz", nil
	}
	if gues > ANumber {
		return "Lütfen daha küçük bir sayı giriniz", nil
	}
	return "Bildinizz", nil
}
func Demo2() {

	mess, hata := guess((79))
	fmt.Println(mess)
	fmt.Println(hata)

}
