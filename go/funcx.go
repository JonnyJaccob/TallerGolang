package main

import (
	"fmt"
)

func calcx(x, y, z int ) int {
	return x + y + z
}


func funcx(){
	fmt.Println(calcx(1,2,3))
}