package main

import "fmt"

type person struct {

FirstName string
LastName string
FavoriteIceCream []string

}


func main(){

	H := person{
		
		FirstName : "H",
		LastName : "f",
		FavoriteIceCream : []string{"a", "b", "c"},

	}


	fmt.Printf("%#v\n", H)
	
	for _, v := range H.FavoriteIceCream {

		fmt.Println(v)
	}
}
