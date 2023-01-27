package main

import "fmt"

/*
func name() {
	statements
}
*/

func main() {
	result, x = c()
	fmt.Println(result)
	fmt.Println(x)
	r := b(1, 2, 3, 4, 5, 6)
	fmt.Println(r)
}
func a() (int, string) {
	//fmt.Println(12)
	return 122, "wref"
}
func b(args ...int) { //args os a slice
	fmt.Println(args)
}
func c() (i int, j string){
	i = 10
	j = "rehul"
	return

}
