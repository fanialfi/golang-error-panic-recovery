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