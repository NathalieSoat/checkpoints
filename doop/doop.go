package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func Power(n int, m int) int {
	helper := n
	if m == 0 {
		n = 1
	} else {
		for i := 1; i < m; i++ {
			n *= helper
		}
	}
	return n
}

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	stri := []rune(s)

	res := 0
	offset := 0

	if stri[0] == 45 || stri[0] == 43 {
		offset = 1
	} else {
		offset = 0
	}
	for i := offset; i < len(s); i++ {

		if '0' <= stri[i] && stri[i] <= '9' {
			exp := len(s) - 1 - i
			if stri[i] == '0' {
				res += 0 * Power(10, exp)
			} else if stri[i] == '1' {
				res += 1 * Power(10, exp)
			} else if stri[i] == '2' {
				res += 2 * Power(10, exp)
			} else if stri[i] == '3' {
				res += 3 * Power(10, exp)
			} else if stri[i] == '4' {
				res += 4 * Power(10, exp)
			} else if stri[i] == '5' {
				res += 5 * Power(10, exp)
			} else if stri[i] == '6' {
				res += 6 * Power(10, exp)
			} else if stri[i] == '7' {
				res += 7 * Power(10, exp)
			} else if stri[i] == '8' {
				res += 8 * Power(10, exp)
			} else if stri[i] == '9' {
				res += 9 * Power(10, exp)
			}
		} else {
			return 0
		}
	}

	if stri[0] == 45 {
		res *= -1
	}
	return res
}

func IsNumeric(s string) bool {
	f := []rune(s)
	len := 0

	for range s {
		len++
	}
	for i := 0; i < len; i++ {
		if f[i] < 48 || f[i] > 57 {
			return false
		}
	}
	return true
}

func main() {

	arr := os.Args

	lengthArgs := len(arr)

	if lengthArgs != 4 {
		return
	}

	if !IsNumeric(arr[1]) {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	if !IsNumeric(arr[3]) {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	if arr[2] == "+" {

		fmt.Print(Atoi(arr[1]) + Atoi(arr[3]))
		z01.PrintRune('\n')

	} else if arr[2] == "-" {

		fmt.Print(Atoi(arr[1]) - Atoi(arr[3]))
		z01.PrintRune('\n')

	} else if arr[2] == "*" {

		fmt.Print(Atoi(arr[1]) * Atoi(arr[3]))
		z01.PrintRune('\n')

	} else if arr[2] == "/" {

		if arr[3] == "0" {
			fmt.Print("No division by 0")
			z01.PrintRune('\n')
			return
		} else {
			fmt.Print(Atoi(arr[1]) / Atoi(arr[3]))
			z01.PrintRune('\n')
		}

	} else if arr[2] == "%" {

		if arr[3] == "0" {
			fmt.Print("No Modulo by 0")
			z01.PrintRune('\n')
			return
		} else {
			fmt.Print(Atoi(arr[1]) % Atoi(arr[3]))
			z01.PrintRune('\n')
		}

	} else {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
}
