package main

import "fmt"

func main() {
	//slice
	 numbers:=[]int{0,5,9,10}
	//los indices de los slices comienzan en 0
	//numbers:=[]int{5,9,10}

	//count es length(numbers)
	//iterar en  los slice
	/*for i := 0; i < len(numbers); i++ {
	fmt.Println(numbers[i])}*/

	//opcion #2 para iterar en los slice
	/*for i := range numbers{
		fmt.Println(numbers[i])
	}*/

	//range devuelve el indice y el valor
	/*for i,number := range numbers{
		fmt.Println("Index: ",i, "Value: ",number)
	}*/

	//omitir el indice
	/*for _,number := range numbers{
		fmt.Println("Value: ",number)
	}*/

	//numbers= append(numbers, 8,3)
	//fmt.Println(numbers)

	//numbers=numbers[inicial:final]
	//toma en cuenta el 1ero, el ultimo es excluyente
	//numbers=numbers[1:5]
	//fmt.Println(numbers)

	//toma el primero y el ultimo
	//numbers=numbers[1:4]
	//fmt.Println(numbers)

	//obtener todos los valores excepto el ultimo
	numbers=numbers[0:4]
	fmt.Println(numbers)

	//es lo mismo que con 0:4
	//numbers=numbers[:4]
	//fmt.Println(numbers)

	//obtenemos desde el 3ero hasta el final
	//numbers = numbers[3:]
	//fmt.Println(numbers)

	//obtnenemos desde el 2do hasta el final
	//numbers = numbers[2:]
	//fmt.Println(numbers)

	//obtenemos banano y las demas son omitidas
	/*fruits:=[]string{"Manzana","Banano","Sandia","Melon"}
	  for _,fruit :=range fruits{
			/*if fruit=="Banano"{
				fmt.Println("Fruit: ",fruit)
			}else{
				continue
			}


	    if fruit != "Banano"{
				continue
			}
			fmt.Println("Fruit: ",fruit)
		}
	*/

	//omitir la fruta si no termina en a
	//slice
	/*fruits :=[]string{"Manzana","Banano","Sandia","Melon"}
	for _,fruit := range fruits{
		index:=len(fruit)-1
		letter:= fruit[index:]

		if letter == "a" {
	     continue
		}
	 fmt.Println("Fruit: ",fruit)
	}*/



  //ejercicio de esteban
	/*numeros := make([]int, 0)
	suma := 0
	fmt.Println("Ingresa cuantos numeros deseas sumar: ")
	var cantidad int
	fmt.Scan(&cantidad)

	for i := 0; i < cantidad; i++ {
		fmt.Printf("Ingresa numero a sumar: ")
		var input int
		fmt.Scan(&input)
		numeros = append(numeros, input)
		suma = suma + input
	}
	fmt.Println("el slice es: ", numeros)
	fmt.Println("la suma es: ", suma)

	suma2 := 0
	for _, item := range numeros {
		suma2 = suma2 + item
	}
	fmt.Println("La suma de los numeros del slice es: ", suma2)
	*/



  //pedir por consola y hago el slice
	/*numbers:=make([]int,0)
	fmt.Println("ingresa los numeros por favor: ")
	var cantidad int
	fmt.Scan(&cantidad)

	for i:=0;i<cantidad;i++{
		fmt.Println("numero para identificar: ")
		var input int
		fmt.Scan(&input)
	  numbers=append(numbers,input)

	}
	fmt.Println("el slice es: ",numbers)*/
}




