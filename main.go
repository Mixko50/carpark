package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
)

// โจทย์
// รถจอดฟรี 2 ชั่วโมงแรก หลังจากนั้นจะเสียค่าจอดรถ 100 บาทต่อชั่วโมง
// เมื่อ 22:00 ถึง 10:00 จะเสียค่าปรับ 1000 บาท
// เศษของชั่วโมงถือเป็นชั่วโมงเต็ม

var format = flag.String("mode", "gen", "gen: for generate new parking ticket\ncheck: for check parking fee\n")

func main() {
	flag.Parse()
	godotenv.Load()
	godotenv.Load(".env")

	if *format == "gen" {
		result, err := GenerateParkingTicket()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	} else {
		fmt.Println("Check parking fee")
	}

	//parkTime := time.Date(2019, time.January, 1, 15, 7, 28, 0, time.UTC)
	//leaveTime := time.Date(2019, time.January, 2, 14, 39, 40, 0, time.UTC)
	//fmt.Println(CalculateParkingFee(parkTime, leaveTime))
}
