package main

import "fmt"

//Declarar funciones 1
func imprimirNombre(nombre string){
	fmt.Println("Fuera del main")
	fmt.Println("El nombre es: ",nombre)
}

//Declarar funciones 2
func suma (n1 int,n2 int)int{
		return n1 + n2
	}

//Declarar funciones 3
func resta(n1,n2 int)(r int){
	r=n1-n2
	return
}



func main(){
	//llamar una funcion
	imprimirNombre("EmmaFreezer")
	fmt.Println("Dentro de main")

	fmt.Println(suma(25,66))
	fmt.Println(resta(88,66))

	//firma de funciones
  fmt.Printf("Funcion Suma:   %T ",suma)
	fmt.Printf("Funcion Resta:  %T ",resta)
}
