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

func CalculateParkingFee(parkTime, leaveTime time.Time) (int, error) {
	pricePerHour := 100
	totalPrice := 0
	difference := leaveTime.Sub(parkTime)
	startSuspensionTime := time.Date(parkTime.Year(), parkTime.Month(), parkTime.Day(), 22, 0, 0, 0, time.UTC)

	// Less than 2 hours
	if difference.Minutes() <= freeParkingTime() {
		return 0, nil
	}

	actualMinutes := difference.Minutes() - freeParkingTime()
	if leaveTime.After(startSuspensionTime) {
		totalPrice += 1000
		actualMinutes -= leaveTime.Sub(startSuspensionTime).Minutes()
		if startSuspensionTime.Sub(parkTime).Minutes() > freeParkingTime() {
			totalPrice += CalculateParkingPriceWithTime(actualMinutes, pricePerHour)
		}
	} else {
		totalPrice += CalculateParkingPriceWithTime(actualMinutes, pricePerHour)
	}
	return totalPrice, nil
}

func freeParkingTime() float64 {
	freeTime := time.Hour * 2
	return freeTime.Minutes()
}

func CalculateParkingPriceWithTime(actualMinutes float64, pricePerHour int) int {
	return int(math.Ceil(actualMinutes/60)) * pricePerHour
}
