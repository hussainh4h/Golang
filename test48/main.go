package main

import "fmt"

func main(){
	
	xxs := make([][]string,0)


	fmt.Println(xxs)
	fmt.Println(len(xxs))
	fmt.Println(cap(xxs))

	y := []string{"a", "b", "c"}

	z := []string{"d", "e", "f"}

	xxs= append(xxs, y, z)
	
	fmt.Println("-----------")

	fmt.Println(xxs)
	fmt.Println(len(xxs))
	fmt.Println(cap(xxs))
	fmt.Println("-----------")
	
	for i, s := range xxs {
		fmt.Println(i, s)
		for  j, v := range s {
			fmt.Println(j, v)

		}

	}

}
