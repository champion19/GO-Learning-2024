package main

import "fmt"

//Closure
func print(cadena string){
	fmt.Println(cadena)
}
func print2(cadena string){
	fmt.Println(cadena)
}

func print3(cadena1,cadena2 string){
fmt.Println(cadena1 + cadena2)
}

// func print4(fprint func(string)){
// 	fprint("Hey desde Print4")

// }

func incrementar()func()int{
	i:=0
	return func()(r int) {
		r =i
		i +=2
		return
	}
}

func incremento(){
	i:=0
	i++
	fmt.Println(i)
}


func main(){
cadena:="Hola Mundo"

imprimir:=print

imprimir(cadena)

imprimir2 :=func(){
	fmt.Println(cadena)
}
imprimir2()


imprimir=print2
imprimir("Hey World")

fmt.Printf("Funcion print1: %T", print)
fmt.Printf("Funcion print2: %T", print2)
fmt.Printf("Funcion print3: %T", print3)

//print4(print)

//Las funciones son comparables con nil
var fb func()
if fb ==nil{
	fmt.Println("fb es igual a nil")
}

inc :=incrementar()
fmt.Println("Valor de i: ",inc())
fmt.Println("Valor de i: ",inc())
fmt.Println("Valor de i: ",inc())
fmt.Println("Valor de i: ",inc())

incremento()
incremento()
incremento()

}
