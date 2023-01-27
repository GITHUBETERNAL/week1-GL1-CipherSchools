package main
/*

*/

func loops {
	for i:= 0; i < 3; i++ {
		fmt.Println(i)
	}


	i := 0
	for i < 3 {
		if i == 1 {
			continue
		}
		fmt.Println(i)
		break
	}
	nums := []int {1, 3, 2, 4, 0}
	for _, v := range nums {
		//fmt.Println(k)
		if value == 3 {
			continue
		}
		fmt.Println(v)
	}
	for _, value := range "rahul" {
		fmt.Println(value)
	}
}