package main

import "fmt"

func main(){

	for i := 1; i < 6; i++{

		fmt.Println("this is the outer loop ", i)

		for j := 1; j < 6; j++ {

			fmt.Println("inner loop is working ", j)
		}

	}




}
