package basics

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// Pascal Case
	// for data types for example Structs, Interfaces Enums ect
	// Eg. CalculateArea, UserInfo, NewHttpRequest

	// snake case
	// Eg, user_id, first_name, last_name, http_request
	// for file names constants and variables

	// UPPER CASE
	// exclusively for constants in go  Eg. PI, GRAVITY
	// This convention ensures that constants stand out as being different from variables and their immutability is emphasized.

	// mixedCase
	// Eg. calculateArea, userId, firstName, httpRequest
	//  for variables or certain identifiers especially when dealing with external libraries or adhering to specific naming conventions from other languages or standards.

	// maintain consistency with existing codebase or project standards or conventions across your project or team

	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)
}
