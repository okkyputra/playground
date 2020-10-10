package main

import (
	"fmt"
	"strconv"
)

// define
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int

	// you can do something like firstname, lastname, city string
}

// greeting method (value receiver)
func (p Person) greet() string {
	return "hello, my name is " + p.firstName + " " + p.lastName + " and i am " + strconv.Itoa(p.age) + " old"
}

// hasbirthday method(pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// init person using struct
	person1 := Person{firstName: "Oks", lastName: "aja", city: "jkt", gender: "F", age: 30}
	person2 := Person{firstName: "Okky", lastName: "ganteng", city: "jkt", gender: "M", age: 30}

	// //Alternative
	// person2 := Person{firstName: "Oks1", lastName: "ganteng1", city: "jkt1", gender: "M1", age: 30}

	// fmt.Println(person1, person2)
	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1, person2)

	person1.hasBirthday()
	person1.getMarried("jajal")
	person2.getMarried("coba")
	fmt.Println(person2.greet())
}
