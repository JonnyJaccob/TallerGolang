package main

import (
	"fmt"
)

func calcy(x, y, z int ) (a, b int) {
	a = (x + y + z)
	b = (x + y)
	return 
}


func returnDoble(){
	fmt.Println(calcy(1,2,3))
}