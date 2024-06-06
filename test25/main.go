package main

import (
	"fmt"
	"math/rand"

)

func main(){


	for i := 0; i < 43; i++{
		
		fmt.Println(i)

		x := rand.Intn(5)
	
		switch {
		
		case x == 0 :
			fmt.Println("this a random number zero ", x) 
		case x == 1 :
			
			fmt.Println("this a random number one ", x) 
		case x == 2 :

			fmt.Println("this a random number two ", x) 
		case x == 3 :

			fmt.Println("this a random number three ", x) 
		case x == 4 :
			
			fmt.Println("this a random number four ", x) 
		
		default :

			fmt.Println("this is a random number ", x)
		}

	}
}
