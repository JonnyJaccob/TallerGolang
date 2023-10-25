package main

import (
	"fmt"
)



func main(){
	str := "Hello World"
	switch str{
	case "Hello" + " World":
		fmt.Println("Hello + World!")
	case "Taller":
		fmt.Printf("Este es un taller de go")
	default:
		fmt.Println("Este es un default")
	}
}