package main

import "fmt"

type Engine struct{
  electric bool

}

type Vehicle struct{

  Engine
  make string
  model string
  doors int
  color string

}

func main(){

	car1 := Vehicle{

		Engine: Engine{
			electric : true,
		},
		make: "BMW",
		model: "latest",
		doors: 4,
		color: "Black",
	}
	
	fmt.Printf("%#v\n", car1)

	car2 := Vehicle{

		Engine: Engine{
			electric : true,
		},
		make: "VW",
		model: "olc",
		doors: 4,
		color: "White",
	}

		
	fmt.Printf("%#v\n", car2)


	fmt.Println("car1 make ", car1.make)
	
	fmt.Println("is car2 engine electric? ", car2.electric)
}
