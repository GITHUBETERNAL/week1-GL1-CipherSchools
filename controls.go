package main

import "fmt"

/*
	if else statements
	if(condition) {
		statements
	}

	if(condition) {
		statements
	} else {
		statements
	} 

	if(condition) {
		statements
	} else if {
		statements
	} else {
		statements
	}

	2. switch(data) {
	case 1 : statement
	case 
	}
*/

func ifelse() {
	fmt.Print("Enter A Number: ")
	var input int 
	//input = 10
	fmt.Scanln(&input)

	if input % 2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if input % 2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
	data := 10
	switch (data) {
		case 10 : fmt.Println("Ten")
		fallthrough // execute the next case also
		default : fmt.Println("Not Ten")
	
	}
}

