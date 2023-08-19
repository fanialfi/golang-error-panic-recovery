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

	// contoh penggunaan pembuatan error sendiri dengan function errors.New()
	var name string
	fmt.Print("masukkan nama anda : ")
	fmt.Scanln(&name)

	if valid, err := lib.Validate(name); valid {
		fmt.Println("Hallo", name)
	} else {
		fmt.Println(err.Error())
	}

	// contoh penggunaan pembuatan error sendiri dengan function fmt,Errorf()
	fmt.Print("masukkan password kamu : ")
	fmt.Scanln(&name)

	if valid, err := lib.ValidateLengthPassword(name); valid {
		fmt.Printf("password kamu \"%s\"\n", name)
	} else {
		fmt.Println(err.Error())
	}
}
