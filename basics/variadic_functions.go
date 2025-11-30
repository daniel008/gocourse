package basics

import "fmt"

func VariadicFunctions() {

	// ... Ellipsis operator (...)
	// func functionName(param1 type1, param2 type2, param3...type3) returnType{
	//  code block
	// }

	// func functionName(param1 ...type1) returnType{
	//  code block
	// }

	// fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))

	// statement, total := sum("The sum of 1, 2, 3 is", 1, 2, 3)
	// statement, total := sum("The sum of 1, 2, 3, 4, 5, 6, is", 1, 2, 3, 4, 5, 6)
	// statement, total := sum( 1, 2, 3)
	// statement, total := sum("1")
	// fmt.Println(statement, total)

	numbers := []int{1, 2, 3, 4, 5, 9}

	sequence3, total3 := sum(3, numbers...)

	fmt.Println("Sequence:", sequence3, "Total", total3)
}

// Variadic function's parameter must be the last parameter in the function signature they can not before any other parameters
// func sum(returnString string, nums ...int) (string, int) {
// 	total := 0
// 	for _, v := range nums {
// 		total += v
// 	}
// 	return returnString, total
// }
func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}
