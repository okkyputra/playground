package main

import "fmt"

func main() {
	// arrays
	//var fruitArr [2]string

	//asign values
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"

	//declare and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

}
