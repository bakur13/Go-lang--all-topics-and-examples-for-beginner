package main

func main() {

	//variables.Demo1()
	//--------slices--------------
	//slices.Demo2()
	//slices.Demo4()
	//---------functions-------
	//var sonuc = functions.Demo2(2, 4)
	//println("toplama sonucu", sonuc)
	//functions.Demo1()
	//------------------demo3--------
	// sonuc1, sonuc2, sonuc3, sonuc4 := functions.Demo3(20, 6)
	// fmt.Println("Toplam sonucu:", sonuc1)
	// fmt.Println("Çıkarma sonucu:", sonuc2)
	// fmt.Println("Çarpma sonucu:", sonuc3)
	// fmt.Println("Bölme sonucu:", sonuc4)
	// fmt.Println(functions.Demo4(2, 4, 6))
	// fmt.Println(functions.Demo4(2, 5, 6, 7, 1, 4, 6))
	// numbers := []int{1, 5, 77}
	// fmt.Println(functions.Demo4(numbers...))

	// functions.Demo5()
	// val1, val2 := functions.Demo6()
	// fmt.Println(val1)
	// fmt.Println(val2)
	//--------------MAPS----------------

	//maps.Demo1()
	//-----------------Range--------------------
	//for_range.Demo1()
	//for_range.Demo2()

	//----------------pointers----------

	// number := 25
	// pointers.Demo1(&number)
	// fmt.Println("number of main:", number)

	// numbers := []int{1, 2, 3}
	// pointers.Demo2(numbers)

	// println("numbers of in main", numbers[0])

	//x := 5
	// pointers.Demo3(x)

	// myslc := []int{1, 10, 100}
	// fmt.Println(myslc)

	//pointers.Demo3(myslc)
	// fmt.Println(myslc)
	//----------------structs--------------
	//structs.Demo2()
	//structs.Demo3()
	//----------GO ROUTİNE-------

	// go goroutiness.Oddnumbers()
	// go goroutiness.Doublenumbers()
	// time.Sleep(5 * time.Second)
	// fmt.Println("bitti")
	//goroutiness.Demo2()

	//--------------chanel--------

	// oddnumscn := make(chan int)
	// doublesnumcn := make(chan int)

	// go channels.Oddnumbers(oddnumscn)
	// go channels.Doublenumbers(doublesnumcn)

	// oddnumtotal, doublenumtotal := <-oddnumscn, <-doublesnumcn
	// multi := oddnumtotal * doublenumtotal
	// fmt.Println("Result of multiplication", multi)
	//-------------İNTERFACE--------------
	//interfaces.Demo1()
	//interfaces.Demo2()
	//interfaces.Demo3()
	//-----------DEFER-STATEMENT-----
	//defer_statement.Demo1()
	//-----------ERROR-HANDLİNG-----
	//error_handling.Demo1()
	//error_handling.Demo2()
	//string_fonk.Demo1()
	//-------------RESTFUL-----
	//restful.Demo1()
	//restful.Demo2()
}
