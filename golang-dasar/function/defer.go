package main

import "fmt"

/*
*
defer menunda eksekusi hingga fungsi selesai.
Berguna untuk cleanup otomatis (misal: file.Close()).
Dieksekusi dalam urutan LIFO.
Membantu menangani panic dengan recover().
*/
func main() {
	runApplicatin()
}

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApplicatin() {
	defer logging() // keyword : defer (function logging akan di jalankan setelah aplikasi selesai di jalankan)
	fmt.Println("Menjalankan Aplikasi")
}
