package main

import "fmt"

func main(){
	var resultado bool
	resultado = 5 < 6
	fmt.Println("5 < 6 - ",resultado)
	resultado=(5 < 6) && (3 > 1)
	fmt.Println("(5 < 6) && (3 > 1) - ", resultado)
	resultado= (5 > 6) || (3 > 1)
	fmt.Println(("5 > 6) || (3 > 1) - "),resultado)
	fmt.Println("!resultado",!resultado)
}
