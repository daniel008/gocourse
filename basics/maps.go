package basics

import (
	"fmt"
	"maps"
)

func Maps() {
	// var mapVariable map[keyType]valueType
	// mapVariable = make(map[keyType]valueType])

	// using map literal
	// mapVariable := map[keyType]valueType{key1: value1, key2: value2}
	// mapVariable := map[string]string{"key1": "value1", "key2": "value2"}
	// fmt.Println(mapVariable)

	myMap := make(map[string]int)
	fmt.Println("Empty map:", myMap)

	myMap["key1"] = 9
	myMap["code"] = 18
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"])
	myMap["code"] = 25
	fmt.Println(myMap)

	// deleting a key-value pair using delete built-in function
	delete(myMap, "key1")
	fmt.Println(myMap)

	// clearing a map using clear built-in function
	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	fmt.Println(myMap)
	// clear(myMap)
	// fmt.Println("After clearing:", myMap)

	// value, unknownvalue := myMap["key1"]
	_, ok := myMap["key1"]
	if ok {
		fmt.Println("A value exists with key1")
	} else {
		fmt.Println("No value exists with key1")
	}
	// fmt.Println(value)
	fmt.Println("Is a value associated with key1:", ok)

	myMap2 := map[string]int{"a": 1, "b": 2, "c": 3}
	myMap3 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(myMap2)
	if maps.Equal(myMap2, myMap3) {
		fmt.Println("myMap2 and myMap3 are equal")
	}

	for k, v := range myMap3 {
		fmt.Println(k, v)
	}

	for _, v := range myMap3 {
		fmt.Println(v)
	}

	// this is a nil map
	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("The map is empty or uninitialized.")
	} else {
		fmt.Println("The map is not empty.")
	}
	val := myMap4["key"]
	fmt.Println(val)

	// this will cause a runtime panic as we are trying to add a key-value pair to a nil map
	// myMap4["key"] = "value"
	// fmt.Println(myMap4)

	// instead, we should initialize the map first
	myMap4 = make(map[string]string)
	myMap4["key"] = "value"
	fmt.Println(myMap4)

	fmt.Println("myMap4 length is", len(myMap4))
	fmt.Println("myMap4 length is", len(myMap))

	myMap5 := make(map[string]map[string]string)
	myMap6 := make(map[string]string)
	myMap6["key1"] = "value1"
	myMap6["key2"] = "value2"

	myMap5["map1"] = myMap4
	fmt.Println(myMap5)
	myMap5["map2"] = myMap6
	fmt.Println(myMap5)
}
