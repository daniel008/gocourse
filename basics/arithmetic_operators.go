package basics

import (
	"fmt"
	"math"
)

func arithmeticOperators() {
	// variable declaration
	var a, b int = 10, 20
	var result int

	result = a + b
	println("Addition: ", result)

	result = a - b
	println("Subtraction: ", result)

	result = a * b
	println("Multiplication: ", result)

	result = a / b
	println("Division: ", result)

	result = a % b
	println("Modulus: ", result)

	const p float64 = 22.0 / 7.0
	fmt.Println(p)

	// overflow with signed integer
	var maxInt = 9223372036854775807 // max value of int64
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	// overflow with unsigned integer
	var uMaxInt uint64 = 18446744073709551615 // max value of uint64
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64

	fmt.Println(smallFloat)
}
