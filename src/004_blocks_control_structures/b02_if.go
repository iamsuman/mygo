package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(5)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
}
