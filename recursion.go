package main

import (
	"fmt"
)

func main() {
	fmt.Println(fact(3))
	rec(3)
}
func rec(num int) {
	if num == 0 {
	return
	}
	rec (num - 1)
	rec (num - 2)
	fmt.Println(num -1)
	if num % 2 == 0 {
		fmt.Println(num + 1)
		rec(num - 1)
		} else {
		fmt.Println(num + 2 )
	}
	fmt.Println(num -1 )
}
func fact(number int) int {
	if number == 1 || number == 0 {
		return 1
	}

	if number < 0 {
		fmt.Println("Invald Number")
		return -1
	}
	result := number * fact(number-1)
	return result
}

func fib(number int) int {
	result := fib(number-1) + fib(number-2)
	return result
}
