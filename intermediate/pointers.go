package intermediate

import "fmt"

func Pointers() {

	var ptr *int
	var a int = 10
	ptr = &a // referencing. & is used to get the address of a variable

	fmt.Println(a)
	fmt.Println(ptr)
	// fmt.Println(*ptr) // dereferencing a pointer
	// if ptr == nil {
	// 	fmt.Println("Pointer is nil")
	// }

	modifyValue(ptr)
	fmt.Println(a)
}

func modifyValue(ptr *int) {
	*ptr++
}
