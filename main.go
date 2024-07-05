package main

import (
	"fmt"
	"math"
	"time"
)

// โจทย์
// รถจอดฟรี 2 ชั่วโมงแรก หลังจากนั้นจะเสียค่าจอดรถ 100 บาทต่อชั่วโมง
// เมื่อ 22:00 ถึง 10:00 จะเสียค่าปรับ 1000 บาท
// เศษของชั่วโมงถือเป็นชั่วโมงเต็ม

func main() {
	fmt.Println("Hello, playground")
	parkTime := time.Date(2019, time.January, 1, 10, 0, 0, 0, time.UTC)
	leaveTime := time.Date(2019, time.January, 1, 13, 0, 0, 0, time.UTC)
	println(CalculateParkingFee(parkTime, leaveTime))
}

func CalculateParkingFee(parkTime, leaveTime time.Time) int {
	pricePerHour := 100
	difference := leaveTime.Sub(parkTime)

	// Less than 2 hours
	if difference.Minutes() <= 120 {
		return 0
	}

	actualMinutes := difference.Minutes() - 120
	return int(math.Ceil(actualMinutes/60)) * pricePerHour
}
