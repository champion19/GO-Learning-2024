package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var h int
var n int
var count int

func iteracion(m int) {

	//fmt.Println("ingresa un numero por favor: ")
	//fmt.Scan(&m)
	h = m

	for k := 0; k < m; k++ {

		for i := 1; i < h; i++ {
			fmt.Print(" ")

		}
		count++

		for j := 0; j < count; j++ {
			fmt.Print("#")
		}
		h--
		fmt.Println()
	}
}

func main() {
	//fmt.Println("ingresa un numero por favor: ")
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int(nTemp)
	iteracion(m)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
