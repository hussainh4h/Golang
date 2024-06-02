package main


import "fmt"


func main(){

	a := [5]int{}

	/*
	a[0] = 1

	a[1] = 2

	a[2] = 3

	a[3] = 4
	
	a[4] = 5
*/


	for i := 0; i < 5; i++ {
		a[i] = i

	}
	
	for _, v := range a {
		fmt.Println(v)
		fmt.Printf("the value is : %v, and the type is: %T\n", v, v)
	
	}
}
