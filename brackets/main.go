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
	if len(os.Args) == 1 {
		z01.PrintRune('\n')
		return
	}

	for i := 1; i < len(os.Args); i++ {
		arr := []rune(os.Args[i])
		var stack []rune

		if os.Args[i] == "" {
			PrintStr("OK")
			z01.PrintRune('\n')
		} else {
			for j := 0; j < len(os.Args[i]); j++ {
				if arr[j] == '(' || arr[j] == '[' || arr[j] == '{' {
					stack = append(stack, arr[j])
				} else {
					switch arr[j] {
					case ')':
						if stack[len(stack)-1] == '(' {
							stack = stack[0 : len(stack)-1]

						}
					case ']':
						if stack[len(stack)-1] == '[' {
							stack = stack[0 : len(stack)-1]

						}
					case '}':
						if stack[len(stack)-1] == '{' {
							stack = stack[0 : len(stack)-1]
						}
					}
				}
			}

			if len(stack) == 0 {
				PrintStr("OK")
				z01.PrintRune('\n')
			} else {
				PrintStr("Error")
				z01.PrintRune('\n')
			}

		}
	}
}
