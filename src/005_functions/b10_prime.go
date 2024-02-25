package main

import (
	"errors"
	"fmt"
)

func get_prime(num int) (result bool, err error) {
	fmt.Println(num)
	if num <= 0 {
		result = false
		return result, errors.New("Number should be greater than 0")
	}
	if num == 1 {
		result = false
		return result, errors.New("Number 1 is a composite numeber")
	}
	result = true
	for i := 2; i < num/2+1; i++ {
		if num%i == 0 {
			result = false
			break
		}
	}
	return result, err
}

func main() {
	fmt.Println(get_prime(1))
	fmt.Println(get_prime(2))
	fmt.Println(get_prime(4))
	fmt.Println(get_prime(10))
	fmt.Println(get_prime(11))
	fmt.Println(get_prime(56))
}
