package main

import( "fmt"
"unsafe"
)

func main(){
// Enteros con signo
//var entero8 int8
//var entero16 int16
var entero32 int32
var entero64 int64

//Enteros sin signo
//var uentero8 uint8
//var uentero16 uint16
//var uentero32 uint32
//var uentero64 uint64

//alias
//var enteroByte byte // alias para uint8
var enteroRune rune // alias para int32

// tipos dependientes
//var enteroUint uint // 32 a 64 bits
var enteroInt int   // 32 a 64 bits
//var enteroUintptr uintptr //

entero32=25
entero64=85
fmt.Println(entero32 + int32(entero64))

//suma int32 y rune
enteroRune=46
fmt.Println(entero32 + enteroRune)

enteroInt=56
fmt.Println(entero32 + int32(enteroInt))

fmt.Println(unsafe.Sizeof(enteroInt),unsafe.Sizeof(entero32))

}






