package main

import "fmt"

func main(){
/*
	var dia int
	fmt.Println("Digita el numero del dia de la semana")
	fmt.Scanln(&dia)

	if dia== 1{
		fmt.Println("lunes")
	}else if dia ==2{
		fmt.Println("martes")
	}else if dia ==3{
		fmt.Println("miercoles")
	}else if dia ==4{
		fmt.Println("jueves")
	}else if dia ==5{
		fmt.Println("viernes")
	}else if dia ==6{
		fmt.Println("sabado")
	}else if dia ==7{
		fmt.Println("domingo")
	}else{
		fmt.Println("Numero invalido")
	}
	fmt.Println("Fin del programa")

switch dia {
case 1:
	fmt.Println("lunes")
case 2:
	fmt.Println("martes")
case 3:
	fmt.Println("miercoles")
case 4:
	fmt.Println("jueves")
case 5:
	fmt.Println("viernes")
case 6:
	fmt.Println("sabado")
case 7:
	fmt.Println("domingo")
default:
	fmt.Println("numero invalido")

}
*/


numero:=26
switch{
case numero > 23:
	fmt.Println("Es mayor que 23")
	fallthrough
case numero > 25:
	fmt.Println("Es mayor que 25")
default:
	fmt.Println("al menos un numero XD")
}

}
