package main

import "fmt"

// able to pass function as a parameter is called closure
func main() {
	number := 10
	fmt.Println(number)
	// store a function as a value to a variable
	var getInt func(int) int
	getInt = func(x int) int {
		fmt.Println("In a function")
		return 20 + x
	}
	getInt(1)
	j := func(x int) int {
		fmt.Println("In a function")
		return 20 + x
	}(19)
	fmt.Println(j)
}

func g(getInt func(int) int) {
	getInt(78)
}

// function is a first class member in golang
