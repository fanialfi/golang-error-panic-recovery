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

## penggunaan function `recover()`

recover digunakan untuk meng-handle panic error. 
Pada saat panic error muncul, recover men-take over (mengambil alih) goruntine yang sedang panic (pesan panic tidak akan muncul)

contoh :

```go
package main

import (
  "fmt"
  "strings"
  "errors"
)

func Catch() {
	// jika nilai kembalian recover() tidak sama dengan nill (terjadi error)
  // untuk menggunakan recover, anonymous function/ function/ closure function dimana recover() berada, harus di eksekusi dengan cara di defer
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func Validate(input string) (bool, error) {
	// function TrimSpace akan menghapus karakter spasi sebelum dan sesudah string
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

func main(){
  // seperti yang sudah dijelaskan di atas, anonymous function/ function/ closure function dimana recover() berada, harus di eksekusi dengan cara di defer
  defer Catch()

  var name string
  fmt.Printf("masukkan nama anda : ")
  fmt.Scanln(&name)

  if valid, err := Validate(name); valid {
    fmt.Println("Hallo", name)
  } else {
    panic(err.Error())
    fmt.Println("END")
  }
}
```

pada saat program di atas dijalankan ketika variabel `name` kosong, maka tidak akan memunculkan panic error seperti yang sudah didelkarasikan di statement `return false, errors.New("cannot be empty")` pada function `Validate`, melainkan panic error tersebut di handle oleh `Catch()` function yang telah di-defer sebelumnya.

dan `recover()` hanya bisa menangkap panic error di goruntine itu sendiri, karena panic error terjadi di goruntine main, dan `recover()` dihandle di function `Catch()` yang telah didefer sebelumnya.
sama seperti behaviour panic error, ketika ada statement yang di defer, maka statement yang didefer akan tetap di eksekusi, sebelum program benar benar menghentikan eksekusinya, baru panic error ditampilkan.

contoh penerapan `recover()` pada function iife :

```go
package main

import "fmt"

func main(){
  defer func(){
    if r := recover(); r != nil {
      fmt.Println("terjadi error dengan pesan :", r)
    } else {
      fmt.Println("tidak ada error")
    }
  }()

  panic("Some error happen")
}
```

contoh penerapan `recover()` pada perulangan, umumnya jika terjadi panic error maka proses setelahnya akan terhenti, mengakibatkan perulangan juga terhenti secara paksa, pada contoh berikut akan diterapkan cara handle panic error tanpa menghentikan program itu sendiri

```go
package main

import "fmt"

func main(){
  hero := []string{"superman", "aquaman", "wonder woman"}

  for _ item := range hero {
    func(){
      defer func(){
        if r := recover(); r != nil{
          fmt.Println("pannic occured  on lopping", item, "| message", r)
        } else {
          fmt.Pritln("application running perfect")
        }
      }()
      panic(fmt.Errorf("some error happend on looping %s", item))
    }()
  }
}
```

pada code di atas didalam sebuah perulangan terdapat sebuah function iife untuk recover panic dan juga kode untuk men-triger panic error secara paksa, ketika panic error terjadi, maka idealnya perulangan terhenti. 
Tapi pada contoh di atas tidak, karena semua operasi sudah di bungkus kedalam iife, dan karena sifat panic error akan menghentikan proses block code yang ada di block function.