package basics

import "fmt"

func Defer() {
	process(10)

}
func process(i int) {
	defer fmt.Println("Deferred i value:", i)
	defer fmt.Println("First deferred statement executed")
	defer fmt.Println("Second deferred statement executed")
	defer fmt.Println("Third deferred statement executed") // going to execute in reverse order
	i++
	fmt.Println("Normal statement executed")
	fmt.Println("i value:", i)
}
