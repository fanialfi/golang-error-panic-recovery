package lib

import (
	"fmt"
	"time"
)

func SendMsg(ch chan<- string) {
	for i := 0; i < 5; i++ {
		now := time.Now()
		h, m, s := now.Clock()

		ch <- fmt.Sprintf("sekarang jam %d:%d:%d:%d", h, m, s, now.Nanosecond())
	}
}

func ReceiveMessage(ch <-chan string) {
	for data := range ch {

		// anonmous function untuk membungkus statement defer
		func() {
			// recovery untuk anonymous function
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("terjadi panic error dalam melooping", data, "| message :", r)
				} else {
					fmt.Println("Nothing Wrong")
				}
			}()
			fmt.Println(data)
		}()
	}
}

/*
	ketika program di atas dipanggil di function main()
	msg := make(chan string, 3)
	go lib.SendMsg(msg)
	lib.ReceiveMessage(msg)

	maka kana terjadi deadlock,
	karena akan terjadi 2 proses yang saling tunggu,
	Deadlock adalah situasi di mana dua atau lebih goroutine dalam program saling menunggu satu sama lain untuk melepaskan sumber daya yang mereka butuhkan. Ini mengakibatkan program terjebak dan tidak dapat melanjutkan eksekusi lebih lanjut. Saya akan menjelaskan lebih detail mengenai deadlock dan mengapa hal itu bisa terjadi dalam kode yang Anda berikan.

Dalam kode yang Anda berikan, deadlock dapat terjadi karena ada dua goroutine yang saling menunggu satu sama lain. Mari kita perinci langkah-langkah yang terjadi:

1. Anda membuat channel `msg` dan menjalankan goroutine dengan `SendMsg(msg)`.
2. Fungsi `SendMsg` berjalan di goroutine terpisah dan mencoba mengirim data melalui channel `msg`.
3. Pada saat `SendMsg` berjalan, channel tersebut menjadi penuh karena tidak ada yang membaca dari channel.
4. Kemudian, Anda memanggil fungsi `ReceiveMessage(msg)` dalam goroutine utama (main).
5. Fungsi `ReceiveMessage` berjalan di goroutine terpisah dan mencoba membaca dari channel `msg`.
6. Namun, channel tersebut tidak akan pernah dikosongkan karena `SendMsg` masih berjalan dan mengirim data ke channel.

Sebagai hasil dari langkah-langkah di atas, kedua goroutine sekarang berada dalam situasi deadlock: `SendMsg` menunggu agar data dapat ditulis ke channel, sedangkan `ReceiveMessage` menunggu agar data dapat dibaca dari channel. Karena keduanya saling menunggu satu sama lain, program akan terjebak dan tidak akan melanjutkan eksekusi lebih lanjut.

Untuk menghindari deadlock dalam situasi ini, Anda dapat menggunakan mekanisme yang disebut buffering pada channel atau menggunakan teknik koordinasi seperti WaitGroup. Buffering channel akan memungkinkan goroutine untuk terus berjalan tanpa harus menunggu saat ada sedikit delay dalam proses pembacaan atau penulisan channel. Sedangkan WaitGroup adalah cara untuk menunggu sejumlah goroutine selesai sebelum melanjutkan eksekusi program.

Dalam contoh kode yang telah saya berikan dalam respons sebelumnya, penggunaan WaitGroup membantu menghindari deadlock dengan memastikan bahwa goroutine `main` menunggu hingga semua pekerjaan selesai sebelum program berakhir.
*/
