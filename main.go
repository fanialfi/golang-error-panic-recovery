package main

import (
	"fmt"
	"golang/error-panic-recovery/lib"
	"os"
)

func main() {
	var input string

	// contoh sederhana penggunaan error
	fmt.Print("Type some number : ")
	fmt.Scanln(&input)

	num, err := lib.KonversiInt(input)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	fmt.Println(num)
}
