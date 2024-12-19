package main

import "fmt"

//Call stack
//LIFO last input first output
func f1(){
	fmt.Println("Entrando a F1")
	f2()
	fmt.Println("Saliendo de F1")
}
func f2(){
	fmt.Println("Entrando a F2")
	f3()
	fmt.Println("Saliendo de F2")
}
func f3(){
	fmt.Println("Entrando a F3")

	fmt.Println("Saliendo de F3")

}

func main(){
	fmt.Println("Entrando en Main")
	f1()
	fmt.Println("Saliendo de Main")
}
