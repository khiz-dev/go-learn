package exercises

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func Exercise3() {
	fmt.Println("Exercise 3 -- ")

	emp1 := Employee{"John", "Doe", 101}
	emp2 := Employee{firstName: "Jane", lastName: "Smith", id: 102}
	emp3 := struct {
		firstName string
		lastName  string
		id        int
	}{
		firstName: "Alice",
		lastName:  "Johnson",
		id:        103,
	}

	fmt.Printf("Employee 1: %+v\n", emp1)
	fmt.Printf("Employee 2: %+v\n", emp2)
	fmt.Printf("Employee 3: %+v\n", emp3)
}
