package main

import (
	"fmt"
	"time"
)

func main() {

	//waktu sekarang
	waktu := time.Now()
	fmt.Println(waktu.Local())

	//definisikan waktu
	utc := time.Date(2024, time.December, 19, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	/**
	Year
	Year: "2006" "06"
	Month: "Jan" "January" "01" "1"
	Day of the week: "Mon" "Monday"
	Day of the month: "2" "_2" "02"
	Day of the year: "__2" "002"
	Hour: "15" "3" "03" (PM or AM)
	Minute: "4" "04"
	Second: "5" "05"
	AM/PM mark: "PM"
	*/

	/**
	time zone
	"-0700"     ±hhmm
	"-07:00"    ±hh:mm
	"-07"       ±hh
	"-070000"   ±hhmmss
	"-07:00:00" ±hh:mm:ss
	*/
	formatWaktu := "2006-01-02 15:4:05"

	value := "2024-12-19 10:10:10"
	valueTime, err := time.Parse(formatWaktu, value)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valueTime)
	}
	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())

	//duration
	fmt.Println(valueTime.Hour())

}
