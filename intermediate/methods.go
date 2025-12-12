package intermediate

import "fmt"

type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width

}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor // r.length = r.length * factor
	r.width *= factor
}
func Methods() {
	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()
	fmt.Println("Area of rectangle with width 9 and length 10 is", area)
	rect.Scale(2)
	area = rect.Area()
	fmt.Println("Area of rectangle after scaling by factor 2 is", area)

	num := MyInt(-5)
	num1 := MyInt(9)
	fmt.Println("Is num positive?", num.IsPositive())
	fmt.Println("Is num1 positive?", num1.IsPositive())
	fmt.Println(num.welcomeMessage())

	shape := Shape{Rectangle{length: 10, width: 9}}
	fmt.Println("Area of shape is", shape.Area())
	fmt.Println("Area of shape is", shape.Rectangle.Area())
}

type MyInt int

// Method on a user-defined type
func (m MyInt) IsPositive() bool {
	return m > 0
}

func (MyInt) welcomeMessage() string {
	return "Welcome to MyInt type in Go!"
}
