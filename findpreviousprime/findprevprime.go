package main

import "fmt"

func FindPrevPrime(nb int) int {
	result := nb
	if nb < 2 {
		return 2
	}
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return FindPrevPrime((nb - 1))
		}
	}
	return result
}

func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
}
