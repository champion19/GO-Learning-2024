package main

import (
	"fmt"
	"strconv"
)

func main(){

	var cadena string
	cadena="soy cineasta"
	fmt.Println(cadena)
	fmt.Println(len(cadena))
	fmt.Println(cadena[2])
	fmt.Println(cadena[0:4])


cadena= cadena+ " nuevamente"
fmt.Println(cadena)

cadena += " siiii "
fmt.Println(cadena)

cadena = `
<html>
<head>
<meta charset="utf-8">
<title></title>
</head>
<body>
</body>
</html>
`
fmt.Println(cadena)


edad := 27
cadena = "la edad es " + strconv.Itoa(edad)
fmt.Println(cadena)
fmt.Println("Edad",edad)


}
