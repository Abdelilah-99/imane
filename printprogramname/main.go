package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argsWithProg := os.Args[0]
	var arr []rune
	for i := 0; i < len(argsWithProg); i++ {
		arr = append(arr, rune(argsWithProg[i]))
	}
	for i := 2; i < len(arr); i++ {
		z01.PrintRune(arr[i])
	}
	z01.PrintRune('\n')
}
