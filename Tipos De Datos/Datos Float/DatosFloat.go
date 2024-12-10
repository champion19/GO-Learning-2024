package main

import "fmt"

func main() {

	var f32 float32     //6 digitos decimales de precision
	var f64 float64     //15 """""" """"""""" "" """"""""
	var c64 complex64   //numero complejo para float32
	var c128 complex128 //"""" """""""" """" float64

	fmt.Println("f32,f64,c64,c128 = ", f32, f64, c64, c128)
	f32 = 0.156
	f64 = 0.156
	fmt.Println("f32 * 58.565656 = ", f32*58.565656)
	fmt.Println("f64 * 58.565656 = ", f64*58.565656)
	c64 = complex(5, 6)
	fmt.Println("c64 = ", c64)
	fmt.Println("c64 * 85458.65 = ", c64*85458.65)
	c128 = complex(5, 6)
	fmt.Println("c128 = ", c128)
	fmt.Println("c128 * 85458.65 = ", c128*85458.65)

}
