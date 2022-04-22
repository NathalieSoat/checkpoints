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
	count := 0
	d, _ := strconv.Atoi(os.Args[1])
	l := append(PrimeFactors(d))
	for i := 0; i < len(l); i++ {
		count++
		if count >= i && i > 0 {
			PrintStr("*")
		}
		PrintStr(strconv.Itoa(l[i]))

	}
	z01.PrintRune('\n')

}

func PrimeFactors(x int) []int {
	factors := []int{}

	candidate := 2
	for candidate <= x {
		if x%candidate == 0 {
			factors = append(factors, candidate)
			x = x / candidate
		} else {
			candidate++
		}
	}

	return factors
}
