package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// myName := person{}
	// fmt.Println(myName)
	myName := person{
		firstName: "Kiet",
		lastName:  "Nguyen",
		contactInfo: contactInfo{
			email:   "ntgkiet@gmail.com",
			zipCode: 12,
		},
	}

	myName.updataName("asdjfkl")
	myName.print()
}

func (p *person) updataName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
