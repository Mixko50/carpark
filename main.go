package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
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

	if os.Getenv("PARKING_SECRET") == "" {
		log.Fatal("PARKING_SECRET is not set")
	}

	if *format == "gen" {
		result, err := GenerateParkingTicket()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	} else {
		fmt.Println(GetParkingFee("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbnRyeV90aW1lIjoxNzIwMjk2MDAwLCJpYXQiOjE3MjAzNjI3NzV9.V8PTPjmRmAw5ga4dX0I2q8ub-YsNmUf5buvFkcL9iZo"))
	}

	//parkTime := time.Date(2019, time.January, 1, 15, 7, 28, 0, time.UTC)
	//leaveTime := time.Date(2019, time.January, 2, 14, 39, 40, 0, time.UTC)
	//fmt.Println(CalculateParkingFee(parkTime, leaveTime))
}
