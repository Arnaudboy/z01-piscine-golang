package main

import (
	"os"

	"github.com/01-edu/z01"
)

// this func print the sqare
func QuadA(x, y int) {
	yMax := &y //this is usefull to stop the loop
	xMax := &x
	if x <= 0 || y <= 0 { // Check for valid input
		return
	}

	for line := y; line >= 1; line-- {
		if line == *yMax || line == 1 {
			PrintLine('o', '-', xMax, yMax)
		} else {
			PrintLine('|', ' ', xMax, yMax)
		}
	}
}

func PrintLine(a rune, b rune, x *int, y *int) { // Avoid repetition in QuadA, this print a line of the square
	for char := *x; char >= 1; char-- {
		if char == *x || char == 1 {
			z01.PrintRune(a)
		} else {
			z01.PrintRune(b)
		}
	}
	z01.PrintRune('\n')
}

func main() {
	//parsing arguments and use it in QuadA func
	if len(os.Args) < 3 || len(os.Args) > 3 || Atoi(os.Args[2]) <= 0 || Atoi(os.Args[1]) <= 0 {
		os.Exit(1)
	} else {
		QuadA(Atoi(os.Args[1]), Atoi(os.Args[2]))
	}
}

//Atoi takes a string and converts it to an integer.
func Atoi(s string) int {
	var res int
	var k int
	var vrai bool
	var Z int
	for a, _ := range s {
		k = a
	}
	if s == "" {
		return 0
	}
	if s[0] == 45 {
		res *= -1
		vrai = true
	}
	if s[0] == 45 || s[0] == 43 {
		Z = 1
	} else {
		Z = 0
	}
	for v := Z; v <= k; v++ {
		if s[v] >= 48 && s[v] < 58 {
			res *= 10
			res += int(s[v] - 48)
		} else if s[v] == 48 {
		} else {
			return 0
		}
	}
	if vrai {
		res = -res
	}
	return res
}
