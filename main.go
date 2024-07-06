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
	//parkTime := time.Date(2019, time.January, 1, 20, 0, 0, 0, time.UTC)
	//leaveTime := time.Date(2019, time.January, 3, 21, 0, 0, 0, time.UTC)
	parkTime := time.Date(2019, time.January, 1, 20, 0, 0, 0, time.UTC)
	leaveTime := time.Date(2019, time.January, 5, 23, 0, 0, 0, time.UTC)
	fmt.Println(CalculateParkingFee(parkTime, leaveTime))
}

func CalculateParkingFee(parkTime, leaveTime time.Time) (int, error) {
	pricePerHour := 100
	totalPrice := 0
	difference := leaveTime.Sub(parkTime)

	// Less than 2 hours
	if difference.Minutes() <= freeParkingTime() {
		return 0, nil
	}

	// Cut the free parking time
	actualMinutes := difference.Minutes() - freeParkingTime()

	if leaveTime.Sub(parkTime).Hours() > 24 {
		// Find day difference
		dayDifference := int(leaveTime.Sub(parkTime).Hours() / 24)

		// Calculate price for each day
		// 2200 is the price for 24 hours
		// 200 is the discount for first 2 hours
		totalPrice += (2200 * dayDifference) - 200

		// Set the park time to be the day as leave time
		parkTime = parkTime.Add(time.Hour * 24 * time.Duration(dayDifference))

		// Calculate the actual minutes for the last day
		actualMinutes = math.Abs(leaveTime.Sub(parkTime).Minutes())

		if isSuspensionTime(leaveTime) {
			totalPrice += 1000

			// Cut the suspension time
			actualMinutes -= leaveTime.Sub(time.Date(leaveTime.Year(), leaveTime.Month(), leaveTime.Day(), 22, 0, 0, 0, time.UTC)).Minutes()
		}

		totalPrice += CalculateParkingPriceWithTime(actualMinutes, pricePerHour)
		return totalPrice, nil
	}

	// Only one day
	startSuspensionTime := time.Date(parkTime.Year(), parkTime.Month(), parkTime.Day(), 22, 0, 0, 0, time.UTC)
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

func isSuspensionTime(leaveTime time.Time) bool {
	return leaveTime.Hour() >= 22 || leaveTime.Hour() < 10
}
