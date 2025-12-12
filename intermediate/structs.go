package intermediate

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell
}

type PhoneHomeCell struct {
	home string
	cell string
}

type Address struct {
	city    string
	country string
}

func (a Person) fullName() string {
	return a.firstName + " " + a.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

func Structs() {

	p := Person{"John", "Doe", 30, Address{"London", "U.K"}, PhoneHomeCell{"01234567", "0123456789"}}

	p1 := Person{
		firstName: "Jane",
		age:       25,
	}

	p2 := Person{firstName: "Jane", age: 25}

	// p1.address.city = "New York"
	// p1.address.country = "U.S.A"

	fmt.Println(p, p1)
	fmt.Println(p.firstName)
	fmt.Println(p1.firstName)
	fmt.Println(p.fullName())
	fmt.Println(p1.fullName())
	fmt.Println("Before increment", p.age)
	p.incrementAgeByOne()
	fmt.Println("After increment", p.age)
	// fmt.Println("Before increment", p1.age)
	// p1.incrementAgeByOne()
	// fmt.Println("After increment", p1.age)
	fmt.Println(p.address.country)
	fmt.Println(p1.address.country)
	fmt.Println(p.cell)
	fmt.Println(p1.address.city)
	fmt.Println(p1, p2)
	fmt.Println("Are p2 and p1 equal?", p2 == p1)

	// Anonymous Structs
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "pseudoemail@example.com",
	}
	fmt.Println(user)
}
