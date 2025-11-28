package basics

import (
	"fmt"
	"slices"
)

func Slices() {

	// var sliceName[]ElementType

	// var numbers []int
	// var numbers1 = []int{1, 2, 3}

	// numbers2 := []int{7, 8, 9}

	// slice := make([]int, 5)

	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4]
	fmt.Println(slice1)

	slice1 = append(slice1, 6, 7)
	fmt.Println("append a slice:", slice1)

	// sliceCopy := make([]int, 5)
	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println("Slice Copy:", sliceCopy)

	var nilSlice []int
	fmt.Println("Is nilSlice nil?", nilSlice)

	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	fmt.Println("Element at index 3 of slice1 is:", slice1[3])

	// slice1[3] = 50
	// fmt.Println("Element at index 3 of slice1 is:", slice1[3])

	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slice1 and sliceCopy are equal")
	}

	// multidimensional slice
	twoDSlice := make([][]int, 5)
	for i := 0; i < 5; i++ {
		innerLen := i + 1
		twoDSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDSlice[i][j] = i + j
			fmt.Printf("Adding value %d in outer slice at index %d, and in inner slice index of %d\n", i+j, i, j)
		}
	}

	fmt.Println("Two Dimensional Slice:", twoDSlice)

	slice2 := slice1[2:4]
	fmt.Println("slice2:", slice2)

	fmt.Println("Capacity of slice2:", cap(slice2))
	fmt.Println("The length of slice2:", len(slice2))
}
