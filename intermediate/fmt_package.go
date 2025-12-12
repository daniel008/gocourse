package intermediate

import "fmt"

func Fmtpackage() {

	// // Printing Actions
	// fmt.Print("Hello")
	// fmt.Print("World! ")
	// fmt.Print(12, 456)

	// fmt.Println("Hello")
	// fmt.Println("World")
	// fmt.Println(12, 456)

	// name := "John"
	// age := 25
	// fmt.Printf("My name is %s and I am %d\n years old", name, age)
	// fmt.Printf("My name is %b and I am %X\n years old\n", age, age)

	// // Formatting Functions
	// s := fmt.Sprint("Hello", "World!", 123, 456)
	// fmt.Println(s)

	// s = fmt.Sprintln("Hello", "World!", 123, 456)
	// fmt.Print(s)
	// fmt.Print(s)

	// sf := fmt.Sprintf("My name is %s and I am %d years old", name, age)
	// fmt.Println(sf)
	// fmt.Print(sf)

	// Scanning Functions
	// var name string
	// var age int

	// fmt.Print("Enter your name and age: ")
	// // fmt.Scan(&name, &age)
	// // fmt.Scanln(&name, &age)
	// fmt.Scanf("%s %d", &name, &age)
	// fmt.Println("Name:", name, "Age:", age)
	// fmt.Printf("Name: %s Age: %d", name, age)

	// Error Formatting Functions
	canDrive, err := checkAge(19)
	if err != nil {
		fmt.Println("Error:", err)
	} else if canDrive {
		fmt.Println("You are able to drive")
	}

}

func checkAge(age int) (bool, error) {
	if age < 18 {
		return false, fmt.Errorf("Age %d is too young to drive.", age)
	}
	return true, nil
}
