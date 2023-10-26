package main

import (
	"fmt"
	"math"
)
type Vertex struct{
	X int
	Y int
}

var (
	v1 = Vertex{1, 2} //has type Vertex
	v2 = Vertex{X: 1} // Y:0 es implicito
	v3 = Vertex{} // X:0 y Y:0 es implicito
	p = &v1 // has type *Vertex
)

func main(){
	fmt.Println("________")
	//struct
	q := *p //*Vertex
	p.X = 2
	fmt.Println(v1,p,v2,v3,q)

	//arrays 
	var a [2]string 
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0],a[1])
	fmt.Println(a)

	primes := [6]int {2,3,5,7,11,13}
	fmt.Println(primes)

	//slices - lista aprox - fragmento de arreglo
	var sl []int = primes[1:4]
	fmt.Println(sl)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a1 := names[0:2]
	b1 := names[1:3]
	fmt.Println(a1,b1)  

	b1[0] = "XXX"
	fmt.Println(a1,b1)  
	fmt.Println(names)  
	
	//slice literal 
	sli := []struct{
		i int 
		b bool
	}{
		{2,true},
		{3,false},
	}
	fmt.Println(sli)

	// For the array 
	//var a [10]int 
	// these slice expressions are equivalent:
	//a[0:10]
	//a[:10]
	//a[0:]
	//a[:]

	s := []int{2,3,5,7,9,11,13}
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	//range 

	//metodos 
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	//interfaces
	var iy interface{}
	describe(iy)

	iy = 42
	describe(iy)

	iy = "Hello"
	describe(iy)

	var i I = T{"hello"}
	i.M()

	// funciones como valores 
	hypot := func(x,y float64) float64{
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5,12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func printSlice( s []int){
	fmt.Printf("len=%d cap=%d %v\n",len(s),cap(s),s)
}

func (v Vertex) Abs() float64{
	//v equivalente a this
	return math.Sqrt(float64(v.X)*float64(v.X) + float64(v.Y) * float64(v.Y))
}

type I interface{
	M()
}

type T struct{
	s string
}

func (t T) M() {
	fmt.Println(t.s)
}

func compute(fn func(float64, float64) float64) float64{
	return fn(3,4)
}
func describe(i interface{}){
	fmt.Printf("(%v)\n", i)
}