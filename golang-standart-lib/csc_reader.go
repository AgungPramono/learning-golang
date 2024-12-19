package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {

	csvData := "Nama,Umur,Alamat\n" +
		"Andi,25,Jakarta\n" +
		"Budi,30,Bandung\n" +
		"Citra,28,Surabaya\n" +
		"Dewi,22,Yogyakarta\n" +
		"Eko,35,Semarang\n" +
		"Fahri,27,Malang\n" +
		"Gita,26,Medan\n" +
		"Hendra,33,Palembang\n" +
		"Indah,24,Bekasi\n" +
		"Joko,29,Bogor"

	reader := csv.NewReader(strings.NewReader(csvData))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

}
