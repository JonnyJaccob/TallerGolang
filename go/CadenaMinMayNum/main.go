package main

import (
	"fmt"
	"strings"
)
func EsLetraNumero(x byte) (string, int){
	var ch byte = x
	//si es minuscula
	if ch >= 97 && ch <= 122 {
		return "Es una letra minuscula",1
	}else if ch >= 65 && ch <= 90 {
		return "Es una letra mayuscula",2
	}else if ch >= 48 && ch <= 57{
		return "Es un numero",3
	}else {
		return "No es ningun caracter permitido",0
	}
		
}
func main(){
	fmt.Println("________")
	cadena := "H m l z r 1 6 z r /"
	cadena = strings.TrimSpace(cadena)
	entrada := 'H'
	x := byte(entrada)
	fmt.Println(EsLetraNumero(x))
	letras := strings.Split(cadena," ")
	var min, may , num int 
	var tmin, tmay, tnum string 
	for i, str := range letras {
		index := i + 1
		char := str[0]
		carac, val  := EsLetraNumero(char)
		if val == 1 {
			if min != 0 {  tmin += ", "}
			tmin += string(char)
			min ++
		}else if val == 2{
			if may != 0 {  tmay += ", "}
			tmay += string(char) 
			may ++
		}else if val == 3{
			if num != 0 {  tnum += ", "}
			tnum += string(char)
			num ++
		}else{
			index = index +1 -1
		}

        fmt.Printf("caracter num: %v %v \n",index,carac)
    }
	fmt.Printf("Numero de minusculas %v, y fueron: %v\n",min, tmin)
	fmt.Printf("Numero de mayusculas %v, y fueron: %v\n",may, tmay)
	fmt.Printf("Numero de numeros %v, y fueron: %v",num, tnum)
}