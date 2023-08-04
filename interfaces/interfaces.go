package interfaces

import (
	"fmt"
	"math"
)

func MyInterfaces() {
	r := rectangle{myWidth: 3, myHeight: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	myWidth, myHeight float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.myWidth * r.myHeight
}

func (r rectangle) perim() float64 {
	return 2 * r.myWidth * 2 * r.myHeight
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
