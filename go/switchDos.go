package main

import (
	"fmt"
	"time"
)

func swithdos(){
	fmt.Println("When is Friday")
	today := time.Now().Weekday()
	//switch today.Friday{
	switch today{
	case today + 0:
		fmt.Println("Today.")
	default:
		fmt.Println("Too far away")
	}
}
	