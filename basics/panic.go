package basics

import "fmt"

func Painc() {

	// panic(interface{})

	// valid input
	process(10)

	// invalid input
	process(-1)
}

func process(input int) {

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")

	if input < 0 {
		fmt.Println("Before Panic")
		panic("Input must be a non-negative integer")
		// panic will stop the execution of the function
		// any code after panic will not be executed
		fmt.Println("After Panic")
		defer fmt.Println("Deferred 3")
	}
	fmt.Println("Processing input:", input)
}
