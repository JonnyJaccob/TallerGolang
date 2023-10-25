package main

import (
	"fmt"
)

func por(){
	sum := 0
	//
	for i := 0; i < 10; i++ {
		fmt.Printf("i is equal to: %v\n", i)
		sum += i
	}
	fmt.Println(sum)

	//equivalente al while 
	for sum < 1000 {
		fmt.Printf("sum is %v\n", sum)
		sum += sum
	}
	fmt.Println(sum)

	//for{//iteracion infinita}

}