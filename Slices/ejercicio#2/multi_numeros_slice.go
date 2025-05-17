package main
import ("fmt")
func main(){
//ejercicio 2 igual al 1ero pero con multiplicacion
	numbers:=[]int{4,2,10,20,25,30,10,12}
	  resultado:=1
	  for _,number:=range numbers{
	   resultado*=number
	  }
	  fmt.Println("el resultado es: ",resultado)
	}
