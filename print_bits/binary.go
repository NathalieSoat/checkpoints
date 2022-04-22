package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var numOfArgs int = 0
	var binVal [10]int
	var decStrg string
	//var err error/
	var decInt int = 0
	var idx int = 0
	//fmt.Println("The Check Point Exam no 3 Question 5 entered.")
	numOfArgs = len(os.Args)
	idx = 9
	if numOfArgs == 2 {
		decStrg = os.Args[1]
		decInt, _ = strconv.Atoi(decStrg)
		fmt.Println("The Decimal Integer to Convert is:", decInt)
		for decInt > 0 {
			binVal[idx] = decInt % 2
			decInt = decInt / 2
			idx--
		} //for while loop/
		fmt.Println("The Binary of given number is:", binVal)
	} else {
		//Error wrong number of Arguments./
	} //if/
} //main*/
