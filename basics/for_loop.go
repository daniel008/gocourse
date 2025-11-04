package basics

import "fmt"

func forLoop() {

	// for i := 1; i <= 5; i++ {
	// 	println(i)
	// }

	// // iterate over collection
	// numbers := []int{1, 2, 3, 4, 5}
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %v, Value: %v\n", index, value)
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue // continue the loop but skip the rest of lines/statements
	// 	}
	// 	fmt.Println("Odd Number: ", i)
	// 	if i == 5 {
	// 		break // break out of the loop
	// 	}
	// }

	// ASTERISK PATTERN LAYOUT
	// rows := 5

	// // outer loop
	// for i := 1; i <= rows; i++ {
	// 	// inner loop for spaces before stars
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	// inner loop for printing stars
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println() // move to the next line
	// }

	for i := range 10 {
		// fmt.Println(10 - i)
		i++
		fmt.Println(i)
	}
	fmt.Println("We have a lift off!")
}
