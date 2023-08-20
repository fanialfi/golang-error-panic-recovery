package lib

import "fmt"

func Catch() {
	// jika nilai kembalian recover() tidak sama dengan nill (terjadi error)
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

// pemanfaat recover() pada anonymous function
func RecoverAnonym() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("terjadi error dengan pesan :", r)
		} else {
			fmt.Println("tidak ada error")
		}
	}()

	panic("Some error happen")
}
