package main

import (
	"fmt"
)

const PI = 3.14


func constante(){
	const World = "§"
	fmt.Println("Hello", World)
	fmt.Println("Happy",PI,"Day")

	const Thruth = true
	fmt.Println("Go rules?", Thruth)
}