package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func Runes() {
	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!"
	// message2 := "Hello, \rGo!" // Hello -> Go!lo
	message2 := "Hello, Go!"
	rawMessage := `Hello\nGo`
	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("length of message variable is:", len(message))
	fmt.Println("length of message2 variable is:", len(message2))
	fmt.Println("length of rawMessage2 variable is:", len(rawMessage))
	fmt.Println("The first character of message is:", message[0]) // ASCII

	greeting := "Hello "
	name := "Alice"
	fmt.Println(greeting + name)

	// behind the scenes they were all numeric values in computer memory
	str := "apple"           // a has an ASCII value of 97
	str1 := "Apple"          // A has an ASCII value of 65
	str2 := "Banana"         // b has an ASCII value of 98
	str3 := "app"            // a has an ASCII value of 97
	fmt.Println(str1 < str2) // compared by lexicographical order
	fmt.Println(str3 < str1) // compared by lexicographical order
	fmt.Println(str > str1)
	fmt.Println(str > str3)

	for _, char := range message {
		// fmt.Printf("Character at index %d is %c\n", i, char)
		// fmt.Printf("%x\n", char) // hexadecimal value of each character
		fmt.Printf("%v\n", char) // actual ASCII value of each character
	}

	fmt.Println("Rune Count:", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'
	cch := '日'

	fmt.Println(ch)
	fmt.Println(cch)

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", cch)

	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type of ch is %T\n", ch)
	fmt.Printf("Type of cstr is %T\n", cstr)

	const NIHONGO = "日本語" // Japanese text
	fmt.Println(NIHONGO)

	jhello := "こんにちは" // Japanese text "Hello"
	for _, runeValue := range jhello {
		// fmt.Printf("%c\n", runeValue) // こ, ん, に, ち, は actual runes
		fmt.Printf("%v\n", runeValue) // Unicode code point value of each rune in integer format int32
	}

	r := '😊'
	fmt.Printf("%c\n", r)
	fmt.Printf("%v\n", r)

}
