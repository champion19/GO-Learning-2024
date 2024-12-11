package main
import "fmt"


func main(){
/*	number := 1
opcion 1 while

for number <= 5 {
    fmt.Println(number)
    number++
  }
*/
	//multiplication table using while
	multiplier :=1

	 for multiplier <= 10 {

    // producto
    product := 5 * multiplier

    // imprime la tabla en fomrato 5 * 1 = 5
    fmt.Printf("5 * %d = %d\n", multiplier, product)
		//%d es un especificador de formato entero de una tabla de multiplicar
    multiplier++
  }


//do while
/*
number := 1

  for {
    if number > 5 {
      break;
    }
    fmt.Printf("%d\n", number);
    number ++
  }
	*/
}
