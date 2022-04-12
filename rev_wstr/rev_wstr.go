package main

import "fmt"

func RevSplitWhiteSpaces(s string) []string {
	k := 0
	bool := false
	for index := range s {
		if bool && index != 0 && (s[index-1] == '\n' || s[index-1] == '\t' || s[index-1] == ' ') && s[index] != '\n' && s[index] != '\t' && s[index] != ' ' {
			k++
		}
		if s[index] != '\n' && s[index] != '\t' && s[index] != ' ' {
			bool = true
		}

	}
	k++

	x := 0
	result := make([]string, k)
	ok := true
	for _, value := range s {
		if value == '\n' || value == '\t' || value == ' ' {
			if !ok {
				x++
			}
			ok = true
			continue
		}
		result[x] = result[x] + string(value)
		ok = false
	}
	l := len(result)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func main() {

	fmt.Println((RevSplitWhiteSpaces("This is a test")))
}
