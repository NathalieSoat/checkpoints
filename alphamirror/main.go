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
		if f[i] >= 'a' && f[i] <= 'm' || f[i] >= 'A' && f[i] <= 'M' {
			for j := 'a'; j <= 'm'; j++ {
				if f[i] == j {
					f[i] += 25 - ((j - 'a') * 2)
				}
			}
			for j := 'A'; j <= 'M'; j++ {
				if f[i] == j {
					f[i] += 25 - ((j - 'A') * 2)
				}
			}
		} else if f[i] >= 'n' && f[i] <= 'z' || f[i] >= 'N' && f[i] <= 'Z' {
			for j := 'n'; j <= 'z'; j++ {
				if f[i] == j {
					f[i] = f[i] - 1 - ((j - 'n') * 2)
				}
			}
			for j := 'N'; j <= 'Z'; j++ {
				if f[i] == j {
					f[i] = f[i] - 1 - ((j - 'N') * 2)
				}
			}
		}
	}
	PrintStr(string(f))
	z01.PrintRune('\n')
}
