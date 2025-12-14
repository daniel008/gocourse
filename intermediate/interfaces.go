package intermediate

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type rect1 struct {
	width, height float64
}

func (r rect1) area() float64 {
	return r.width * r.height
}
func (r rect1) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func Interfaces() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	r1 := rect1{width: 3, height: 4}

	measure(r)
	measure(c)
	measure(r1)

	myPrinter(1, "John", 3.14, true)
	printType(1)
	printType("John")
	printType(3.14)
	printType(true)
	printType([]int{1, 2, 3})
}

func printType(i interface{}) {
	switch i := i.(type) {
	case int:
		fmt.Println("I'm an int", i)
	case string:
		fmt.Println("I'm a string", i)
	case float64:
		fmt.Println("I'm a float64", i)
	case bool:
		fmt.Println("I'm a bool", i)
	default:
		fmt.Println("I don't know what I am", i)
	}
}

func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}
