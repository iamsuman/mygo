package main

import "fmt"

func main() {
	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	p := person{
		FirstName: "Pat",
		//   MiddleName: "Perry", // This line won't compile
		MiddleName: stringp("Perry"), // This works
		LastName:   "Peterson",
	}
	fmt.Println(*p.MiddleName)
}

func stringp(s string) *string {
	return &s
}
