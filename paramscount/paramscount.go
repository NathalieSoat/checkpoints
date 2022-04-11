package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lenght := len(os.Args[1:])
	z01.PrintRune(rune(lenght + 48))
	z01.PrintRune('\n')
}
