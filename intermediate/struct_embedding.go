package intermediate

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	employeeInfo person // Embedding person struct with a field name employeeInfo which is of type person when we access fields or methods of person struct we need to use employeeInfo field to access fields or methods of person struct
	person              // Embedding person struct anonymously when we access fields or methods of person struct we can access directly without using any field name
	emId         string
	salary       float64
}

func (p person) introduce() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Printf("Hello, my name is %s and I am %d years old. I am an employee with ID %s and salary %.2f.\n", e.employeeInfo.name, e.employeeInfo.age, e.emId, e.salary)
}

func StructEmbedding() {
	emp := Employee{
		employeeInfo: person{name: "Alice", age: 30},
		emId:         "E123",
		salary:       75000,
	}

	fmt.Println("Name", emp.employeeInfo.name) // Accessing embedded struct field directly we don't need to use emp.person.name
	fmt.Println("Age", emp.employeeInfo.age)
	fmt.Println("Emp ID", emp.emId)
	fmt.Println("Salary", emp.salary)

	emp.introduce() // Calling method of embedded struct directly we don't need to use emp.person.introduce()
}
