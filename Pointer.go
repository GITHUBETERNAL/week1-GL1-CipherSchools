package main
import "fmt"
func main() {
	// i := 10
	// j := &i 
	// *j = 100
	var i int
	i = 10 
	// var j *int 
	j := &i 
	*j = 100 
	j = new(int) 
	fmt.Println(j)
	fmt.Println(i)

	name := new(string) 
	*name = "golang"
	fmt.Println(*name)

}