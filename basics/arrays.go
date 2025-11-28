package basics

import "fmt"

func Arrays() {
	// var arrayName [size]elementType
	var numbers [5]int
	fmt.Println("Array:", numbers)

	numbers[4] = 20
	fmt.Println("Updated Array:", numbers)

	numbers[0] = 9
	fmt.Println("First Element:", numbers)

	fruits := [4]string{"Apple", "Banana", "Orange", "Mango"}
	fmt.Println("Fruits Array:", fruits)
	fmt.Println("Second Fruit:", fruits[1])

	// originalArray := [3]int{1, 2, 3}

	// copiedArray := originalArray
	// copiedArray[0] = 100

	// fmt.Println("Original Array:", originalArray)
	// fmt.Println("Copied Array:", copiedArray)

	for i := 0; i < len(fruits); i++ {
		fmt.Println("Fruit at index", i, "is", fruits[i])
	}

	for i, v := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", i, v)
	}

	// underscore is blank identifier, used to store unused values
	for _, v := range fruits {
		fmt.Printf("Fruit: %s\n", v)
	}

	a, _ := someFunction()
	fmt.Println(a)
	// fmt.Println(b)

	b := 2
	_ = b

	fmt.Println("The length of fruits array is:", len(fruits))

	// Comparing arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	fmt.Println("Are arrays equal?", array1 == array2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("2D Array (Matrix):", matrix)

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	copiedArray[0] = 100

	fmt.Println("Original Array:", originalArray)
	fmt.Println("Copied Array:", copiedArray)
}

func someFunction() (int, int) {
	return 1, 2
}
