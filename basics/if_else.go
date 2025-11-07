package basics

import "fmt"

func ifElse() {
	//  if condition {
	//  block of code
	// }

	// age := 25

	// if age >= 18 {
	// 	fmt.Println("You are old enough to vote")
	// }

	//  if condition {
	//  block of code
	// } else if {
	//  block of code
	// } else {
	//  block of code
	// }

	// temperature := 25
	// if temperature >= 30 {
	// 	fmt.Println("It's hot outside!")
	// } else {
	// 	fmt.Println("It's not too hot outside.")
	// }

	// score := 95
	// if score >= 90 {
	// 	fmt.Println("You got an A!")
	// } else if score >= 80 {
	// 	fmt.Println("You got a B.")
	// } else if score >= 70 {
	// 	fmt.Println("You got a C.")
	// } else {
	// 	fmt.Println("You failed the exam.")
	// }

	// if score >= 90 {
	// 	fmt.Println("You got an A!")
	// }
	// if score >= 80 {
	// 	fmt.Println("You got a B.")
	// }
	// if score >= 70 {
	// 	fmt.Println("You got a C.")
	// } else {
	// 	fmt.Println("You failed the exam.")
	// }

	// if condition1 {
	// 	block of code1
	// 	if condition2 {
	// 		code block2
	// 	}
	// }

	num := 15
	// if num%2 == 0 {
	// 	if num%3 == 0 {
	// 		fmt.Println("The number is divisible by both 2 and 3.")
	// 	} else {
	// 		fmt.Println("The number is divisible by 2 but not by 3.")
	// 	}
	// } else {
	// 	fmt.Println("The number is not divisible by 2.")
	// }

	// || - logical OR
	// && - logical AND
	if num%2 == 0 && num%3 == 0 {
		fmt.Println("The number is divisible by both 2 and 3.")
	} else if num%2 == 0 || num%3 == 0 {
		if num%2 == 0 {
			fmt.Println("The number is divisible by 2.")
		}

		fmt.Println("The number is divisible by 3.")

	} else {
		fmt.Println("The number is not divisible by 2 or 3.")
	}

}
