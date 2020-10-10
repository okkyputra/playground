package main

import "fmt"

func main() {

	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	//use * for read val from address
	fmt.Println(a, *b)

	//change val with pointer
	*b = 10
	fmt.Println(a)
}
