package basics

import "fmt"

func switchCase() {

	// switch statement in go is (switch case default) (fallthrough keyword is used to continue with the next case)
	// switch expression {
	// case value1:
	// 	// code to be executed if expression is equal to value1
	// fallthrough
	// case value2:
	// 	// code to be executed if expression is equal to value2
	// case value3:
	// 	// code to be executed if expression is equal to value]]]]
	// default:
	// 	// code to be executed if none of the above cases match
	// 	println("default")
	// }

	// switch statement  in other languages (switch case break default)
	// switch expression {
	// case value1:
	// 	// code to be executed if expression is equal to value1
	// break;
	// case value2:
	// 	// code to be executed if expression is equal to value2
	// break;
	// case value3:
	// 	// code to be executed if expression is equal to value3
	// break;
	// default:
	// 	// code to be executed if none of the above cases match
	// 	println("default")
	// }

	// fruit := "watermelon"
	// switch fruit {
	// case "apple":
	// 	fmt.Println("Apple")
	// case "banana":
	// 	fmt.Println("Banana")
	// case "orange":
	// 	fmt.Println("Orange")
	// default:
	// 	fmt.Println("Invalid fruit")
	// }

	// Multiple Conditions
	// day := "Monday"

	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a Weekday.")
	// case "Saturday", "Sunday":
	// 	fmt.Println("It's a Weekend!")
	// default:
	// 	fmt.Println("Invalid day.")
	// }

	// Switch with expression
	// number := 15

	// switch {
	// case number < 10:
	// 	fmt.Println("Number is less than 10.")
	// case number > 10 && number < 20:
	// 	fmt.Println("Number is between 10 and 19.")
	// default:
	// 	fmt.Println("Number is greater than 20 .")
	// }

	// Switch with fallthrough
	// num := 6
	// switch {
	// case num > 1:
	// 	fmt.Println("Greater than 1")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("Equal to 2")
	// default:
	// 	fmt.Println("not 2")
	// }
	checkType(5)
	checkType(3.14)
	checkType("Hello")
	checkType(true)

}
func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Integer")
		// fallthrough can not fallthrough to another data type by default
	case int32:
		fmt.Println("Integer 32")
	case float64:
		fmt.Println("Float")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Unknown Type")
	}

}
