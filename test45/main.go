package main

import "fmt"


func main(){

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

//	x[11] = 52

	fmt.Println(len(x))

	fmt.Println(cap(x))
	//fmt.Println(x)

	x = append(x, 52)
	fmt.Println("after appending 52", x)

	x = append(x, 53, 54, 55)
	fmt.Println("after appending the 3 numbers", x)

	y := []int{56, 57, 58, 59, 60}

	x = append(x, y...)
	fmt.Println("after appending the y slice", x)

	fmt.Println(len(x))


	fmt.Println(cap(x))


}
