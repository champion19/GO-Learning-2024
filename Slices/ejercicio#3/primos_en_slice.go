package main
	import("fmt")
 //nuevo desafio leer e imprimir los numeros primos de un slice
//funcion normal
func esPrimo (n int) bool {
	if n<2{
		return false
	}
	for i:=2; i<n-1;i++{
		if n%i==0{
			return false
		}
	}
	return true
 }

func main(){
 numbers:=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13}
	primos:=[]int{}//slice para guardar los primos
//funcion anonima
/*esPrimo:=func (n int) bool {
	if n<2{
		return false
	}
	for i:=2; i<n-1;i++{
		if n%i==0{
			return false
		}
	}
	return true
 }
*/
 for _,v:= range(numbers){
		if esPrimo(v){
		primos=append(primos,v)
		}
	}
	fmt.Println("los primos son: ",primos)
}



