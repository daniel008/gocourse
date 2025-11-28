package basics

import "fmt"

func Functions() {
	// func <name>(parameters list) returnTypes {
	// code block
	// ...
	// return value
	//}

	// sum := add(1, 2)

	// fmt.Println(sum)
	// 	fmt.Println(add(1, 2))

	// 	greet := func() {
	// 		fmt.Println("Hello, Anonymous Function!")
	// 	}

	// 	greet()

	// 	// operation := add(1, 2)
	// 	// result := operation
	// 	operation := add
	// 	result := operation(3, 5)

	// 	fmt.Println(result)

	// passing function as a parameter
	result := applyOperation(1, 2, add)
	fmt.Println("1 + 2 =", result)

	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 =", multiplyBy2(6))
}

func add(a, b int) int {
	return a + b
}

// This function takes two integers and a function as parameters
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
