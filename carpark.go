package main

import (
	"fmt"
	"math"
	"time"
)

const (
	pricePerHour                    = 100
	freeParkingHours                = 2
	suspensionStartTime             = 22
	suspensionEndTime               = 10
	suspensionFee                   = 1000
	dailyParkingFeeBeforeSuspension = 1200
)

func CalculateParkingFee(parkTime, leaveTime time.Time) int {
	// Round up the seconds to be a minute
	if leaveTime.Second() != 0 {
		leaveTime = leaveTime.Add(time.Minute).Truncate(time.Minute)
	}
	difference := leaveTime.Sub(parkTime)

	// Less than 2 hours
	if difference.Minutes() <= freeParkingTime() {
		return 0
	}

	if leaveTime.Sub(parkTime).Hours() > 24 {
		return CalculateMultipleDayParkingFee2(parkTime, leaveTime)
	}

	return CalculateSingleDayParkingFee(parkTime, leaveTime)
}

func CalculateSingleDayParkingFee(parkTime, leaveTime time.Time) int {
	startSuspensionTime := time.Date(parkTime.Year(), parkTime.Month(), parkTime.Day(), suspensionStartTime, 0, 0, 0, time.UTC)
	endSuspensionTime := time.Date(leaveTime.Year(), leaveTime.Month(), leaveTime.Day(), suspensionEndTime, 0, 0, 0, time.UTC)
	difference := leaveTime.Sub(parkTime)
	actualMinutes := difference.Minutes() - freeParkingTime()
	totalPrice := 0

	if leaveTime.After(startSuspensionTime) {
		totalPrice += suspensionFee
		actualMinutes -= leaveTime.Sub(startSuspensionTime).Minutes()
		if leaveTime.After(endSuspensionTime) {
			actualMinutes += leaveTime.Sub(endSuspensionTime).Minutes()
		}
		if startSuspensionTime.Sub(parkTime).Minutes() > freeParkingTime() {
			totalPrice += CalculateParkingPriceWithTime(actualMinutes, pricePerHour)
		}
	} else {
		totalPrice += CalculateParkingPriceWithTime(actualMinutes, pricePerHour)
	}
	return totalPrice
}

func CalculateMultipleDayParkingFee2(parkTime, leaveTime time.Time) int {
	totalPrice := 0

	startSuspensionTime := time.Date(parkTime.Year(), parkTime.Month(), parkTime.Day(), suspensionStartTime, 0, 0, 0, time.UTC)
	if startSuspensionTime.Sub(parkTime).Minutes() > freeParkingTime() {
		totalPrice += CalculateParkingPriceWithTime(startSuspensionTime.Sub(parkTime).Minutes(), pricePerHour)
	}

	if startSuspensionTime.Sub(parkTime).Minutes() > freeParkingTime() {
		totalPrice -= 200
	}

	dayDifference := int(leaveTime.Sub(parkTime).Hours() / 24)
	if dayDifference > 0 {
		totalPrice += suspensionFee * dayDifference
		totalPrice += dailyParkingFeeBeforeSuspension * (dayDifference - 1)
	}

	endSuspensionTime := time.Date(leaveTime.Year(), leaveTime.Month(), leaveTime.Day(), suspensionEndTime, 0, 0, 0, time.UTC)
	roundUpLeaveTime := leaveTime
	if leaveTime.Minute() != 0 || leaveTime.Second() != 0 {
		roundUpLeaveTime = leaveTime.Add(time.Hour).Truncate(time.Hour)
	}

	if roundUpLeaveTime.Hour() <= suspensionStartTime && endSuspensionTime.Before(roundUpLeaveTime) {
		fmt.Println(roundUpLeaveTime)
		fmt.Println(endSuspensionTime)
		totalPrice += int(math.Abs(endSuspensionTime.Sub(roundUpLeaveTime).Hours())) * pricePerHour
	} else {
		totalPrice += dailyParkingFeeBeforeSuspension
	}

	if isSuspensionTime(leaveTime) {
		totalPrice += suspensionFee
	}

	return totalPrice
}

func freeParkingTime() float64 {
	return time.Hour.Minutes() * freeParkingHours
}

func CalculateParkingPriceWithTime(actualMinutes float64, pricePerHour int) int {
	return int(math.Ceil(actualMinutes/time.Hour.Minutes())) * pricePerHour
}

func isSuspensionTime(leaveTime time.Time) bool {
	return leaveTime.Hour() >= suspensionStartTime || leaveTime.Hour() < suspensionEndTime
}
