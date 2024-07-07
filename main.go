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
var tokenString = flag.String("token", "", "token for check parking fee\n")

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
		if tokenString == nil || *tokenString == "" {
			log.Fatal("Token is required")
		}
		fmt.Println(GetParkingFee(*tokenString))
	}
}
