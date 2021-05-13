package main

import (
	"github.com/01-edu/z01"
	"os"
)

func QuadC(x, y int) {
	yMax := &y
	xMax := &x
	if x <= 0 || y <= 0 {
		return
	}

	for line := y; line >= 1; line-- {
		if line == *yMax {
			PrintLine2('A', 'B', 'A', xMax, yMax)
		} else if line == 1 {
			PrintLine2('C', 'B', 'C', xMax, yMax)
		} else {
			PrintLine2('B', ' ', 'B', xMax, yMax)
		}
	}
}

func PrintLine2(a rune, b rune, c rune, x *int, y *int) {
	for char := *x; char >= 1; char-- {
		if char == *x {
			z01.PrintRune(a)
		} else if char == 1 {
			z01.PrintRune(c)
		} else {
			z01.PrintRune(b)
		}
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) < 3 || len(os.Args) > 3 || Atoi(os.Args[2]) <= 0 || Atoi(os.Args[1]) <= 0 {
		os.Exit(1)
	} else {
		QuadC(Atoi(os.Args[1]), Atoi(os.Args[2]))
	}
}

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
