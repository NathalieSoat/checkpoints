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
	var result []rune

	for i := 0; i <= len(f)-1; i++ {
		if f[i] >= 'a' && f[i] <= 'z' {
			for j := 'a'; j < 'z'; j++ {
				if f[i] == j {
					for k := 1; k <= int(j)-96; k++ {
						result = append(result, j)
					}
				}
			}
		} else if f[i] >= 'A' && f[i] >= 'Z' {
			for j := 'A'; j < 'Z'; j++ {
				if f[i] == j {
					for k := 1; k <= int(j)-64; k++ {
						result = append(result, j)
					}
				}
			}
		} else {
			result = append(result, f[i])
		}
	}
	PrintStr(string(result))
	z01.PrintRune('\n')
}
