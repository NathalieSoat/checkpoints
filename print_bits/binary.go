package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	input := os.Args[0:]
	var rInput string
	var startDiv int = 128
	var divResult int
	var remainder int = 0
	var bin []int
	if len(input) != 2 {
		os.Exit(0)
	} else {
		rInput = os.Args[1]
		step, err := strconv.Atoi(rInput)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Step:", step)
		/*calculation := step * multiplier
		  decimal = decimal + calculation
		  multiplier *= 10
		  } else {
		  for z := 0; z < 8; z++ {
		  z01.PrintRune('0')
		  }
		  os.Exit(0)
		  }
		  }*/
		for startDiv > 1 {
			divResult = step / startDiv
			fmt.Println("Start Div:", startDiv, "Div Result:",
				divResult, "Step:", step, "Remainder:", remainder, "Binary:", bin)
			if divResult == 0 {
				bin = append(bin, 0)
			} else {
				bin = append(bin, 1)
			} /*if*/
			startDiv = startDiv / 2
			divResult = step % startDiv
			remainder = step - startDiv
			divResult = remainder / step
			step -= startDiv
		} /*for loop*/
		fmt.Println("Start Div:", startDiv, "Div Result:",
			divResult, "Binary:", bin)
	}
	for l := len(bin) - 1; l >= 0; l-- {
		z01.PrintRune(rune(bin[l]))
	}
	z01.PrintRune('\n')
}
