package main

import (
	"fmt"
)

func calcz(x, y, z int ) (a, b int) {
	a = (x + y + z)
	b = (x + y)
	return 
}


func flotante(){
	x, y := calcz(1,2,3)
	a, _ := calcz(1,2,3)
	_ , c := calcz(1,2,3)
	fmt.Println(x,y)
	fmt.Printf("type %T, value %v\n", a, a)
	fmt.Printf("type %T, value %v\n", float64(a), float64(a))
	fmt.Printf("type %T, value %v\n", c, c)
	fmt.Printf("type %T, value %v\n", 'c', 'c')
	fmt.Printf("type %T, value %v\n", string('c'), string('c'))
}