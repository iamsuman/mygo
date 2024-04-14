package main

import "fmt"

func main() {
	var x int = 10
	var y int = 127
	var z int
	var z1 int16
	z = x + y
	z1 = int16(z)
	fmt.Println(z, z1)
}
