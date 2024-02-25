package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 1. C style for loop
	fmt.Println("C style for loop")
	for i := 0; i < 10; i++ {
		fmt.Println(i, rand.Intn(10))
	}

	// 2. Condition style for loop
	fmt.Println("Condition style for loop")
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}

	// 3. Infinite for loop
	// for {
	//     fmt.Println("Hello")
	// }

	// 4. for range
	fmt.Println("C style for loop")
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
	// Output
	// 0 2
	// 1 4
	// 2 6
	// 3 8
	// 4 10
	// 5 12

}
