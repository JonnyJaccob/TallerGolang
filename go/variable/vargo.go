package main

import (
		"fmt"
		"math"
)

var c, python, java bool = true, false, true
var cmas bool

func main(){
	fmt.Println(python)
	python := true
	i := 10
	fmt.Println(i,c,java,python,cmas,math.Pi)
}