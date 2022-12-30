package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println(p.productName, "kaydedildi")
	defer log()
}
func log() {
	fmt.Println("loglandı")
}
func Demo1() {
	p := product{productName: "laptop", unitPrice: 5000}
	p = product{productName: "mouse", unitPrice: 45}

	fmt.Println("işlem Başarılı")
	defer p.save()
}
