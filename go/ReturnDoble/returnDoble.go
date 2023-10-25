package main

import (
	"fmt"
)

func calcy(x, y, z int ) (a, b int) {
	a = (x + y + z)
	b = (x + y)
	return 
}


func main(){
	fmt.Println(calcy(1,2,3))
}