package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	pipe()
}

func pipe() {
	stdin := os.Stdin

	// args := os.Args
	// cmd()
	idata, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		printStr("ERROR")
	}
	printStr(string(idata))
	print("\n")
	print(string(os.Stdin))
}

func printStr(str string) {
	for _, c := range str {
		z01.PrintRune(c)
	}
}

// func cmd() {
//     cmd := exec.Command("cat")
//     stdout, _ := cmd.StdinPipe()
//     // fmt.Println(s)
//     fmt.Println(stdout)
//     fmt.Println(cmd)
// }
