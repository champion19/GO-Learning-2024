package main

import "fmt"

func main(){

	var cadena string
	cadena="soy cineasta"
	fmt.Println(cadena)
	fmt.Println(len(cadena))
	fmt.Println(cadena[2])
	fmt.Println(cadena[0:4])


cadena= cadena+ "nuevamente"
fmt.Println(cadena)

}
