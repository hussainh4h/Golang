package main

import "fmt"

func main(){

	m := map[string]int{

		"james": 42,
		"ms lilly": 45,


	}


	m1 := m["james"]

	fmt.Println("James age is ", m1)

	
	for k := range m{

	
		if k == "Q"{
		fmt.Printf("this is the Q value %v\n", k)
		} else {
			fmt.Println("not  found")
		}

	}
	
	fmt.Println("______________________________")
	
	if v, ok := m["Q"]; ok {

		fmt.Println("the value of Q is ", v)

	} else {

		fmt.Println("not  found")

	}

	

}
