package main

import "fmt"


func main(){

	s := []int{}
	
	for i:= 0; i<11;i++{
	s = append(s, i+42)
		fmt.Printf("%v -- %T -- memory address is %p\n", s[i], s[i], &s[i])
		
	}
	

}
