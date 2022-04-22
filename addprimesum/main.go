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

//PrimeNumbers returns an array of prime ints, from 2 to nbr
func PrimeNumbers(nbr int) []int {
	var result []int
	for i := 2; i <= nbr; i++ {
		if IsPrime(i) {
			result = append(result, i)
		}
	}
	return result
}

func IsPrime(nb int) bool {

	if nb <= 0 || nb == 1 {
		return false
	}

	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	result := 0
	d, _ := strconv.Atoi(os.Args[1])
	primes := PrimeNumbers(d)

	for i := 0; i < len(primes); i++ {
		result += primes[i]
	}

	PrintStr(strconv.Itoa(result))
	z01.PrintRune('\n')
}
