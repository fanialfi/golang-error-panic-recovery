# Error, Panic, dan Recovery

`error` merupakan topik yang sangat penting di go, `panic` untuk memunculkan panic error, dan `recovery` untuk mengatasinya.

## pemanfaatan error

`error` merupakan sebuah tipe, error memiliki 1 buah properti berupa method `Error()`, method ini mengembalikan detail pesan error dalam bentuk string, error termasuk tipe yang isinya bisa `nil`.

di go banyak sekali function yang mengembalikan value lebih dari satu, biasanya salahsatu nilai kembaliannya bertipe `error`.
Contohnya seperti `strconv.Atoi()` function tersebut berfungsi untuk mengubah tipe data _string_ menjadi _number_, function ini mengembalikan dua nilai kembalian, nilai pertama adalah nilai hasil konversi dan nilai kedua adalah `error`, ketika konversi berjalan mulus, nilai kedua akan berisi `nil`, sedangkan jika hasil konversi gagal, penyebabnya bisa langsung diketahui dari error yang dikembalikan.

contoh penggunaan :

```go
package main

import (
  "strconv"
  "fmt"
)

func KonversiInt(str string) (int, error) {
	var number int
	var err error

	number, err = strconv.Atoi(str)

	if err == nil {
		return number, nil
	}
	return number, err
}

func main(){
  var input string
  fmt.Print("Type some number : ")
  fmt.Scanln(&input)

  var number int
  var err error

  number, err = KonversiInt(input)
}
```

ketika program di atas dijalankan maka akan mucnul tulisan `Type some number : `, ketik angka bebas, jika sudah maka tekan enter.

statement `fmt.Scanln(&input)` digunakna untuk men-capture inputan yang diketik oleh user, kemudian disimpan divariabel `input`, selanjutnya variabel tersebut dikonversi menggunakan function `strconv.Atoi()`, function tersebut mengembalikan dua data dan ditampung oleh variabel `number` dan `err`.

Data pertama (`number`) berisi hasil konversi, dan data kedua `err` berisi informasi error (jika memang terjadi error saat konversi).

Setelah dilakukan pengecekan, ketika tidak ada error `number` ditampilkan, dan jika ada error, `input` ditampilkan beserta pesan errornya, untuk pesan error bisa dilihat dari method `Error()` milik tipe `error`.

## membuat custom error sendiri

kita juga bisa membuat object error sendiri (membuat custom error) dengan menggunakan function `errors.New()` (harus import package `errors` dulu)

contoh pembuatan custom error, digunakan untuk pengecekan input apakah kosong atau tidak :

```go
package main

import (
  "errors"
  "fmt"
  "strings"
)

func Validate(input string) (bool, error){
  if strings.TrimSpace(input) == "" {
    return false, errors.New("cannot be empty")
  }
  return true, nil
}

func main(){
  var name string
  fmt.Print("masukkan nama anda : ")
  fmt.Scanln(&name)

  if valid, err := Validate(name); valid {
    fmt.Println("Hello", name)
  }else{
    fmt.Println(err.Error())
  }
}
```

function `Validate` diatas mengembalikan 2 data, data pertama adalah boolean yang menandakan apakah inputan valid atau tidak, data kedua adalah pesan error-nya (jika inputan tidak valid A.K.A inputan kosong).

function `strings.TrimSpace()` diatas digunakan untuk menghilangkan karakter spasi di awal dan di akhir string,
ketika inputan tidak valid, error baru dibuat dengan memanfaatkan function `errors.New()`,
selain itu object error bisa juga dibuat dengan menggunakan function `fmt.Errorf()`


contoh penggunaan function `fmt.Errorf()`

```go
package main

import (
  "errors"
  "fmt"
  "strings"
)

func ValidateLengthPassword(input string) (bool, error) {
	data := strings.TrimSpace(input)

	if len(data) <= 8 {
		return false, fmt.Errorf("panjang password harus lebih dari 8\npanjang password kamu : %d dengan value \"%s\"", len(data), data)
	}
	return true, nil
}

func main(){
  var password string

  fmt.Print("masukkan password kamu : ")
	fmt.Scanln(&password)

  	if valid, err := lib.ValidateLengthPassword(name); valid {
		fmt.Printf("password kamu \"%s\"\n", name)
	} else {
		fmt.Println(err.Error())
	}
}
```

function `fmt.Errorf()` menerima minimal 1 parameter, parameter pertama adalah pesan errornya dalam bentuk string, dan setelahnya adalah parameter variadic dengan tipe any, jika parameter pertama mengandung format specifier seperti `%f`, `%d`, `%v`, `%s`, dll, maka parameter setelahnya adalah variabel atau data yang nantinya disimpan dalam format specifier yang telah ditentukan sebelumnya.

Karena function `fmt.Errorf()` mengembalikan value bertipe `error`, maka juga bisa memanggil method `Error()`

## Penggunaan function `panic()`

`panic()` digunakan untuk menampilkan _stack trace error_ sekaligus menghentikan flow goruntine (karena `main()` juga goruntine, maka behaviour yang sama juga berlaku), setelah ada panic maka proses setelahnya akan berhenti kecuali proses yang sudah di-`defer` sebelumnya (akan muncul sebelum panic error), panic menampilkan pesan error sama seperti `fmt.Println()` namun informasi yang ditampilkan adalah _stack trace error_ jadi sangat mendetail.

pada codingan sebelumnya, pada program yang telah dibuat tadi function `fmt.Println()` untuk menampilkan informasi error akan diganti menggunakan `panic()`, pada program baris pertama setelah deklarasi function `main()` tambahkan statement yang di-`defer` dan setelah panic tambahkan statement untuk mencetak sembarang tulisan seperti berikut :

```go
package main

import (
  "errors"
  "fmt"
  "strings"
)

func Validate(input string) (bool, error){
  if strings.TrimSpace(input) == "" {
    return false, errors.New("cannot be empty")
  }
  return true, nil
}

func main(){
  defer fmt.Println("Program berjalan")
  var name string
  fmt.Print("masukkan nama anda : ")
  fmt.Scanln(&name)

  if valid, err := Validate(name); valid {
    fmt.Println("Hello", name)
  }else{
    // fmt.Println(err.Error())
    panic(err.Error())
    fmt.Println("END")
  }
}
```

ketika program di atas dijalankan dan langsung tekan enter maka panic error akan muncul dan baris kode setelahnya tidak dieksekusi tapi statement yang didefer akan muncul sebelum panic error.