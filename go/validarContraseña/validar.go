package main //advertencia no critica

import (
	"fmt"
	"gotaller/CadenaMinMayNum"
)

func main(){
	fmt.Println("________")
	//a,_ := MMNum.EsLetraNumero(x)
	var contra string = "jfajf!1A"
	arregloDeCadenas := make([]string, len(contra))
	for i,char := range contra{
		arregloDeCadenas[i] = string(char)
	}
	min,may,num,_,_,_ := MMNum.FcicloContador(arregloDeCadenas)
	//metodo 1
	if min > 0 && may > 0 && num >0 {
		fmt.Println("Cumple con las condiciones, exito")
	}else{
		fmt.Println("El texto debe tener letras minusculas, mayusculas y numeros")
	}
	//Metodo 2
	if MMNum.FesValido(arregloDeCadenas) {
		fmt.Println("Cumple con las condiciones, exito")
	}else{
		fmt.Println("El texto debe tener letras minusculas, mayusculas y numeros")
	}
}