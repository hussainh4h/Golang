package main

import (
	"fmt"

	"github.com/hussainh4h/puppy"

	"main/greet"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	
	greet.Visiable()
	greet.Internal()
}



