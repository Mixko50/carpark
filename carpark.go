package main

import (
	"math"
	"time"
)

const (
	pricePerHour                    = 100
	freeParkingDiscount             = 200
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

	return CalculateSingleDayParkingFee2(parkTime, leaveTime)
}

func CalculateSingleDayParkingFee2(parkTime, leaveTime time.Time) int {
	totalPrice := 0
	startSuspensionTime := time.Date(parkTime.Year(), parkTime.Month(), parkTime.Day(), suspensionStartTime, 0, 0, 0, time.UTC)
	if leaveTime.Before(startSuspensionTime) {
		totalPrice += CalculateParkingPriceWithTime(leaveTime.Sub(parkTime).Minutes(), pricePerHour)
		if leaveTime.Sub(parkTime).Minutes() > freeParkingTime() {
			totalPrice -= freeParkingDiscount
			return totalPrice
		}
		return totalPrice
	}

	totalPrice += suspensionFee
	if startSuspensionTime.Sub(parkTime).Minutes() > freeParkingTime() {
		totalPrice += CalculateParkingPriceWithTime(startSuspensionTime.Sub(parkTime).Minutes(), pricePerHour)
		totalPrice -= freeParkingDiscount
	}

	endSuspensionTime := time.Date(leaveTime.Year(), leaveTime.Month(), leaveTime.Day(), suspensionEndTime, 0, 0, 0, time.UTC)
	if leaveTime.After(endSuspensionTime) && parkTime.Day() != leaveTime.Day() {
		totalPrice += CalculateParkingPriceWithTime(leaveTime.Sub(endSuspensionTime).Minutes(), pricePerHour)
		return totalPrice
	}

	return totalPrice
}

func CalculateMultipleDayParkingFee2(parkTime, leaveTime time.Time) (totalPrice int) {
	startSuspensionTime := time.Date(parkTime.Year(), parkTime.Month(), parkTime.Day(), suspensionStartTime, 0, 0, 0, time.UTC)
	if startSuspensionTime.Sub(parkTime).Minutes() > freeParkingTime() {
		totalPrice += CalculateParkingPriceWithTime(startSuspensionTime.Sub(parkTime).Minutes(), pricePerHour)
		totalPrice -= freeParkingDiscount
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
