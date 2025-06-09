package main

import (
	"fmt"
)

var arr [5]int
var suma = 0

func Max(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func Min(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func main() {
	arr := []int{1, 3, 5, 7, 9}
	fmt.Println("Arreglo: ", arr)

	mayor := Max(arr)
	menor := Min(arr)
	sumaTotal := 0
	sumaMin := 0
	sumaMax := 0

	for _, val := range arr {
		sumaTotal = sumaTotal + val

		sumaMin = sumaTotal - mayor
		sumaMax = sumaTotal - menor
	}

	fmt.Println("EL mayor es: ", mayor)
	fmt.Println("El menor es: ", menor)
	fmt.Print(sumaMin, sumaMax)
}
