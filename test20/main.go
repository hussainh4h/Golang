package main


import (
	"fmt"
	"math/rand"
)

func main(){
	//generate a random number
	x := rand.Intn(251)
	//print that generated number
        fmt.Println("the value of x is: ",x)
	

/*
        if x >= 0 && x <= 100; {
        x := rand.Intn(251)
                fmt.Println("between 0 and 100")
        } else if x <= 200 {
                fmt.Println("between 101 and 200")
        } else if x <= 250 {
                fmt.Println("between 201 and 250")
        }

*/

	//conditional statement depending on the value of x will start with 0 and its maximum value is 25
	switch {

	case x <= 100:
                fmt.Println("between 0 and 100")
	case x <= 200:
                fmt.Println("between 101 and 200")
	default:
                fmt.Println("between 201 and 250")
	}
	
}

