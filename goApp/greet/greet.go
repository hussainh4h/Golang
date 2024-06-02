package greet

import "fmt"

// create a non visiable function that you can access it only inside the package directory
func notVisiable(){
	fmt.Println("Hello friends")
}


// create a visiable fuc that you can access it outisde the package directory
func Visiable(){
 	fmt.Println("Hello World")

}


