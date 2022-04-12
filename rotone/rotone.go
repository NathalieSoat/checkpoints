package main

import "github.com/01-edu/z01"

func rot1(d rune) rune {
	if d >= 'a' && d <= 'y' || d >= 'A' && d <= 'Y' {
		return d + 1
	}
	if d == 'z' {
		return 'a'
	}
	if d == 'Z' {
		return 'A'
	}
	return d
}

func Rot1(s string) string {
	result := ""
	for _, val := range s {
		result += string(rot1(val))
	}
	return result
}

func main() {
	result := Rot1("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
