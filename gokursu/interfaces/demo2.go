// Polymorphism interface in golang
package interfaces

import "fmt"

type shp interface {
	area() float64
}

func printarea(shapess ...shp) {
	for _, shp := range shapess {
		fmt.Println("Alan:", shp.area())
	}

}

type triangle struct {
	a float64
	h float64
}

func (t triangle) area() float64 {
	return (t.a * t.h) / 2
}

type square struct {
	a float64
}

func (s square) area() float64 {
	return (s.a * s.a)

}

type rectangles struct {
	a, b float64
}

func (r rectangles) area() float64 {
	return (r.a * r.b)

}
func Demo2() {
	t := triangle{225000000000, 759}
	s := square{4}
	r := rectangles{4, 5}
	printarea(t, s, r)
}
