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
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	f := []rune(os.Args[1])
	for i := 0; i <= len(f)-1; i++ {
		if f[i] >= 'a' && f[i] <= 'z' {
			f[i] = f[i] - 32
		}else if f[i] >= 'A' && f[i] <= 'Z' {
			f[i] = f[i] + 32
		}
	}
	PrintStr(string(f))
	z01.PrintRune('\n')
}
