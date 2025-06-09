package main

import "fmt"





func birthdayCakeCandles(candles []int)int {
	max :=candles[0]
	count:=1

	for i:=1;i<len(candles);i++{
   if candles[i]>max{
		max=candles[i]
		count=1
	 }else if candles[i]==max{
		count++
		}
	}

	return count

}

func main(){
	candles := []int{3, 2, 1, 3}
	fmt.Println("Velas más altas que se pueden apagar:", birthdayCakeCandles(candles)) 

}
