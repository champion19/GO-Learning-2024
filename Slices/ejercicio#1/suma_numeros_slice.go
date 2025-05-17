package main
import( "fmt")
func main(){
	// 1er problema de slice de esteban
	numbers:=[]int{10,20,2,3,5,9,90}
	  suma:=0
	  for _,number:= range numbers{
	  	suma +=number
	  	//fmt.Println("La suma es: ",suma)
	  }
	  fmt.Println("la suma es: ",suma)
	  }


