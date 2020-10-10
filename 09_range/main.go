package main

import "fmt"

func main() {
	ids := []int{33, 41, 65, 1, 14, 65}

	// loop trough id
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// adds id together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("sum", sum)

	// range with map
	emails := map[string]string{"oks": "oks@mail", "oks1": "oks1@mail", "oks11": "oks11@mail"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("name: " + k)
	}

}
