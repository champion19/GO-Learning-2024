package main
import "fmt"
var razagoku="Saiyajin"

func main(){
	//nombrar variables
	numero := 25
	_nombre :="Emma"
	numero2 :=54
	nombreusuario := "admin"
  BeastNumber :=666
	fmt.Println(numero,numero2,_nombre,nombreusuario,BeastNumber)
	var(dios="goku"
	enemigo1,enemigo2="freezer","Majin buu"
)
var fecha=1997
fmt.Println(dios,enemigo1,enemigo2,fecha)

//scope

fmt.Println("la raza de guerreros es: " + razagoku)
imprimir()
}

func imprimir(){
	fmt.Println("la raza de goku es: "+ razagoku)
}
