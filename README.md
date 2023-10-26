# Taller de GO

desarrollado por google  
alto rendimiento - lenguaje maquina   

## Comando:
inicializar y darle el nombre
go mod init <nombre>
```shell
go mod init gotaller
```

```shell
go mod tidy
```

```shell
go run .
```

puede usarse :=
en la declaracion de variables, solo en funciones
var i int = 10  
i := 10 

no tiene clases

modulos
funciones 
cruds 
  
return mas de un valor
```go
func calc(x, y, z int ) (int,int) {
	return (x + y + z),(x + y)
}
```
  
break no existe aqui   

defer se ejecuta al final 
igual a funcion asincrona 

## ejercicio
leer cadena y cuales son mayusculs, minusculas y cuales son digitos 

https://github.com/AriAlanPR/gotaller.git

go mod tiny 
go get .