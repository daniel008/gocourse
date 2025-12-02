package basics

import "fmt"

func Recover() {
	processRecover()
	fmt.Println("Returned from processRecover")
}

func processRecover() {
	defer func() {
		// if r := recover(); r != nil { // this is a more concise way
		r := recover()
		if r != nil {
			fmt.Println("Recovered in processRecover", r)
		}
	}()

	fmt.Println("Start ProcessRecover")
	panic("Something went wrong!")
	fmt.Println("End ProcessRecover")
}
