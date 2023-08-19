package lib

import (
	"errors"
	"fmt"
	"strings"
)

// function untuk mengecek apakah inputan dari user kosong atau tidak
// jika kosong kembalikan nilai false, dan object error baru yang dibuat dengan function errors.New()
// jika tidak kembalikan nilai true dan nil
func Validate(input string) (bool, error) {
	// function TrimSpace akan menghapus karakter spasi sebelum dan sesudah string
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

// function untuk mengecek panjang password, apakah sesuai kriteria atau tidak
// jika tidak kembalikan nilai false dan object error baru yang dibuat dengan function fmt.Errorf()
// jika tidak kembalikan nilai true dan nil
func ValidateLengthPassword(input string) (bool, error) {
	data := strings.TrimSpace(input)

	if len(data) <= 8 {
		return false, fmt.Errorf("panjang password harus lebih dari 8\npanjang password kamu : %d dengan value \"%s\"", len(data), data)
	}
	return true, nil
}
