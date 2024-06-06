package main

import "fmt"

func main(){

	m := map[string]int{

		"james": 42,
		"ms lilly": 45,


	}

	for k, v := range m{
		fmt.Printf("the key is %v, the value is %v\n", k, v)
	}




}
