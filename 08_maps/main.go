package main

import "fmt"

func main() {
	// define map
	// emails := make(map[string]string)

	// // assign key value
	// emails["oks"] = "oks@mail.com"
	// emails["oks1"] = "oks1@mail.com"
	// emails["oks11"] = "oks11@mail.com"

	// declare and add key value
	emails := map[string]string{"oks": "oks@mail", "oks1": "oks1@mail"}
	emails["oks11"] = "oks11@mail.com"

	fmt.Println(emails)
	fmt.Println(emails["oks"])

	// delete map
	delete(emails, "oks11")
	fmt.Println(emails)
}
