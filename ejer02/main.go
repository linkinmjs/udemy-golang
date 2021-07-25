package main

import "fmt"

var numero int
var texto string
var status bool = true

func main() {
	numero2, numero3, numero4, texto := 2, 5, 67, "Este es mi texto"

	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	
	mostrarStatus()

}

funct mostrarStatus(){
	fmt.Println(status)
}