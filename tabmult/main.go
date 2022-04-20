package main

import (
	"os"
	"strconv"

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
	for i := 1; i <= 9; i++ {
		PrintStr(strconv.Itoa(i))
		PrintStr(" x ")
		PrintStr(os.Args[1])
		PrintStr(" = ")
		d, _ := strconv.Atoi(os.Args[1])
		result := i * d

		PrintStr(strconv.Itoa(result))
		z01.PrintRune('\n')
	}
}
