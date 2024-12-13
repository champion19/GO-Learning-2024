package main

import "fmt"
//definicion de datos
//custom type
type Profession string
const(
	lawyer Profession="lawyer"
	programmer Profession="Programmer"
)


type Person struct{
	name string
	age int
	Job  Profession
}

func main(){
	//var p Person=Person {name:"Emma",age:27}
	 p :=  Person{name:"Emma",age:27,Job:"Programmer"}
	/*p.name="emma"
	p.age= 27 */
	fmt.Println(p)
}
