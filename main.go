package main

import (
	"fmt"
	"golang/error-panic-recovery/lib"
)

func main() {
	// ketika terjadi panic error, statement yang di defer berikut akan ditampilkan sebelum panic error
	// namun jika ada beberapa statement yang di defer, maka semua statement yang di defer tersebut akan di eksekusi secara terbalik
	// dari yang terakhir di defer akan menjadi yang pertama di defer
	defer fmt.Println("Selamat datang di program saya")

	defer lib.Catch()
	var input string

	// contoh sederhana penggunaan error
	fmt.Print("Type some number : ")
	fmt.Scanln(&input)

	num, err := lib.KonversiInt(input)

	if err != nil {
		// fmt.Println(err.Error())
		panic(err.Error())
		fmt.Println("END 1")
	}
	fmt.Println(num)

	// contoh penggunaan pembuatan error sendiri dengan function errors.New()
	var name string
	fmt.Print("masukkan nama anda : ")
	fmt.Scanln(&name)

	if valid, err := lib.Validate(name); valid {
		fmt.Println("Hallo", name)
	} else {
		// fmt.Println(err.Error())
		panic(err.Error())
		fmt.Println("END 2")
	}

	// contoh penggunaan pembuatan error sendiri dengan function fmt,Errorf()
	fmt.Print("masukkan password kamu : ")
	fmt.Scanln(&name)

	if valid, err := lib.ValidateLengthPassword(name); valid {
		fmt.Printf("password kamu \"%s\"\n", name)
	} else {
		// fmt.Println(err.Error())
		panic(err.Error())
		fmt.Println("END 3")
	}

	// pemanfaatan recover pada anonymous function
	lib.RecoverAnonym()

	fmt.Println()
	hero := []string{"superman", "aquaman", "wonder woman"}

	for _, elm := range hero {

		// function iife A untuk membungkus defer statement
		func() {

			// recovery untuk anonymous function dalam perulangan
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("panic occured on looping", elm, "| message :", r)
				} else {
					fmt.Println("application running perfectly")
				}
			}()

			// jika panic di tempatkan di dalam block function iife A,
			// maka panic error akan di tangkap oleh defer function recover di atas
			panic(fmt.Errorf("some error happend on %s", elm))
		}()
		// jika panic di tempatkan di sini,
		// maka panic error akan di tangkap oleh function Catch
		// panic(fmt.Errorf("some error happend on %s", elm))
		fmt.Println("elm :", elm)
	}

	// proses dibawah ini menyebabkan deadlock
	// mencoba recovery dalam looping menerima data channel
	// fmt.Println()
	// msg := make(chan string, 3)
	// go lib.SendMsg(msg)
	// lib.ReceiveMessage(msg)
}
