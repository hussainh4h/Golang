package main

import "fmt"

func main(){

	s := make([]string, 0, 50)

	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = append(s, `Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `Utah`, ` Vermont`, ` Virginia`, ` Washington`, `West Virginia`, ` Wisconsin`, ` Wyoming`)



	fmt.Println(len(s))
	
	fmt.Println(cap(s))

	fmt.Printf("%#v\n", s)

	for i:=0; i<len(s); i++{
		fmt.Println(i, s[i])


	}

}
