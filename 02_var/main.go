package main

import "fmt"

func main() {
	//variable using var keyword
	//var name = "Okky"
	var age = 30
	var isAlive = true
	//shorthand
	name, email := "Okky", "my@email.com"

	fmt.Println(name, age, email, isAlive)
	fmt.Printf("%T\n", isAlive)
}
