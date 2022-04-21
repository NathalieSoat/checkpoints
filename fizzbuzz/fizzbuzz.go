package main

import "github.com/01-edu/z01"

func PrintStr(str string) {

	for _, letter := range str {
		z01.PrintRune(letter)
	}
}

func PrintNbr(n int) {

	t := 1
	if n < 0 {
		z01.PrintRune('-')
		t = -1
	}
	if n != 0 {
		q := (n / 10) * t
		if q != 0 {
			PrintNbr(q)
		}
		d := ((n % 10) * t) + '0'
		z01.PrintRune(rune(d))
	} else {
		z01.PrintRune('0')
	}
}

func main() {

	for nb := 1; nb <= 100; nb++ {
		if nb%3 == 0 && nb%5 == 0 {
			PrintStr("fizzbuzz")
			z01.PrintRune('\n')
		} else if nb%3 == 0 {
			PrintStr("fizz")
			z01.PrintRune('\n')
		} else if nb%5 == 0 {
			PrintStr("buzz")
			z01.PrintRune('\n')
		} else {
			PrintNbr(nb)
			z01.PrintRune('\n')
		}
	}
}
