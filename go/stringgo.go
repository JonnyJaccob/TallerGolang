package main

import (
	"fmt"
	"strconv"
)

func mystrx1(a int ) string {
	return strconv.Itoa(a)
}


func stringgo(){
	s := "My value is " + mystrx1(45)
	fmt.Printf("type %T, value %v\n", s, s)
}