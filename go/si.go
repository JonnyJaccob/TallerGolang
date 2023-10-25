package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string  {
	a := ""
	if x < 0 {
		a =  sqrt(-x) + "i"
	} else if x = 1; x == 1{
		a =  fmt.Sprint(math.Sqrt(x))
	}else{
		a = "not 1 nor less than 0"
	}
	return a
}

func si(){
	fmt.Println(sqrt(2),sqrt(-4))
	fmt.Println(sqrt(2),sqrt(-5))
}
	