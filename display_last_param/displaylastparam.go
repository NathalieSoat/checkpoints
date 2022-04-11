package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	f := os.Args

	for _, element := range f[len(f)-1] {
		z01.PrintRune(element)
	}
	z01.PrintRune('\n')
}
