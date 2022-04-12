package main

import "github.com/01-edu/z01"

func rot13(d rune) rune {
	if d >= 'a' && d <= 'm' || d >= 'A' && d <= 'M' {
		return d + 13
	}
	if d >= 'n' && d <= 'z' || d >= 'N' && d <= 'Z' {
		return d - 13
	}
	return d
}

func Rot13(s string) string {
	result := ""
	for _, val := range s {
		result += string(rot13(val))
	}
	return result
}

func main() {
	result := Rot13("AkjhZ zLKIJz , 23y ")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
