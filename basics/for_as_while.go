package basics

import "fmt"

func forAsWhile() {

	// i := 1
	// for i <= 5 {
	// 	fmt.Println("Iteration", i)
	// 	i++
	// }

	// Do not run this code it will run forever this is as other languages do while infinite loop.
	// for {
	// 	fmt.Println("Infinite Loop")
	// }

	// for as while with break
	// sum := 0
	// for {
	// 	sum += 10
	// 	fmt.Println("Sum: ", sum)
	// 	if sum >= 50 {
	// 		break
	// 	}
	// }

	num := 1
	for num <= 10 {
		if num%2 == 0 {
			num++
			continue
		}
		fmt.Println("Odd Number", num)
		num++ // increment operator increases value by 1 and decrement operator decreases value by 1.
	}
}
