package basics

import (
	"errors"
	"fmt"
)

func MultipleReturnValues() {

	// func functionName(parameter type1, parameter type2...) (returnType1, returnType2...) {
	//     code block
	//     ...
	//     return value1, value2
	//}

	quotient, remainder := divide(10, 3)
	fmt.Printf("Quotient: %v. Remainder: %v\n", quotient, remainder)

	result, err := compare(10, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}

func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	// return quotient, remainder
	return // implicit return similar to return quotient, remainder, simplify the code
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if a < b {
		return "a is less than b", nil
	} else {
		return "", errors.New("a is equal to b")
	}
}
