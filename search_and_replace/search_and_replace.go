package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(str string) {

	for _, letter := range str {
		z01.PrintRune(letter)
	}
}

func main() {
	if len(os.Args) != 4 {
		z01.PrintRune('\n')
		return
	}
	f := []rune(os.Args[1])
	f2 := []rune(os.Args[2]) // searcheable chr
	f3 := []rune(os.Args[3]) // replaceable chr

	for i := 0; i <= len(f)-1; i++ {
		if f[i] == f2[0] {
			f[i] = f3[0]
		}
	}
	result := ""
	for i := 0; i < len(f); i++ {
		result += string(f[i])
	}
	PrintStr(result)
	z01.PrintRune('\n')
}
