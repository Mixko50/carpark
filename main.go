package main

import (
	"fmt"
	"time"
)

// โจทย์
// รถจอดฟรี 2 ชั่วโมงแรก หลังจากนั้นจะเสียค่าจอดรถ 100 บาทต่อชั่วโมง
// เมื่อ 22:00 ถึง 10:00 จะเสียค่าปรับ 1000 บาท
// เศษของชั่วโมงถือเป็นชั่วโมงเต็ม

func main() {
	fmt.Println("Hello, playground")
	parkTime := time.Date(2019, time.January, 1, 15, 7, 28, 0, time.UTC)
	leaveTime := time.Date(2019, time.January, 2, 14, 39, 40, 0, time.UTC)
	fmt.Println(CalculateParkingFee(parkTime, leaveTime))
}
