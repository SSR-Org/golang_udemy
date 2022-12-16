package main

import "fmt"

// type person struct {
// 	firstName string
// 	lastName  string
// }

// func main() {
// 	// first decleration type
// 	alex := person{"Alex", "Anderson"}

// 	// second decleration type
// 	alex1 := person{lastName: "Anderson", firstName: "Alex1"}
// 	fmt.Println(alex, alex1)

// 	// third decleration type
// 	var alex2 person

// 	alex2.firstName = "Alex2"
// 	alex2.lastName = "Anderson"

// 	fmt.Printf("%+v", alex2)

// 	fmt.Print()
// 	// Print formats using the default formats for its operands and writes to standard output.
// 	// Spaces are added between operands when neither is a string.
// 	// It returns the number of bytes written and any write error encountered.
// 	fmt.Println()
// 	// Println formats using the default formats for its operands and writes to standard output.
// 	// Spaces are always added between operands and a newline is appended.
// 	// It returns the number of bytes written and any write error encountered.
// 	fmt.Printf("")
// 	// Printf formats according to a format specifier and writes to standard output.
// 	// It returns the number of bytes written and any write error encountered.
// }

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Carrey",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 123456,
		},
	}

	// fmt.Printf("%+v", jim)

	jimPointer := &jim
	jim.updateName("jimmy") // wont work, need to pass pointer

	//Pointer Method
	jimPointer.updateNameUsingPointer("JimmyPointer") // works
	jim.print()

	//Pointer Shortcut
	jim.updateNameUsingPointer("JimPointer")
	jim.print()

}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (pointerToPerson *person) updateNameUsingPointer(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
