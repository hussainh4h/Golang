package main

import "fmt"

type Person struct {

FirstName string
LastName string
FavoriteIceCream []string

}


func main(){

	H := Person{
		
		FirstName : "H",
		LastName : "f",
		FavoriteIceCream : []string{"a", "b", "c"},

	}


	A := Person{
		
		FirstName : "A",
		LastName : "F",
		FavoriteIceCream : []string{"D", "E", "F"},

	}


	m := map[string]Person{

		H.LastName : H,
		A.LastName : A,

	}



	//fmt.Println(m)


	fmt.Println("---------------")

	for _, v := range m {
		fmt.Println(v)
		
			
		for _, v2 := range v.FavoriteIceCream {


			fmt.Println(v2)
		}
	}

	fmt.Println("---------------")





}
