package main

import (
	"fmt"
	"math/rand"
)


func main(){

	for i:= 0; i < 100; i++ {
		fmt.Println(i)

	}

	

	for i:= 0; i<100; i++{

		//create 2 random numbers that goes from 0 to 9
	x, y := rand.Intn(10), rand.Intn(10)
	
	fmt.Printf("the value of x is %v, the value of y is %v\n", x, y)

	// create a switch statment to evaluate the 2 random numbers
	
	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y are both less than 4")
		
	case x > 6 && y > 6: 

		fmt.Println("x and y are both greater than 6")
	case x >=4 && x <= 6:


		fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
	case y != 5: 

		fmt.Println("y is not 5")
	default:

		fmt.Println("none of the previous cases were met")
	}

		fmt.Println("-----------new loop------------")
  	}

}