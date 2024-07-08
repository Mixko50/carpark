package main

import (
	"math"
	"time"
)

const (
	PricePerHour                    = 100
	FreeParkingDiscount             = 200
	FreeParkingHours                = 2
	SuspensionStartTime             = 22
	SuspensionEndTime               = 10
	SuspensionFee                   = 1000
	DailyParkingFeeBeforeSuspension = 1200
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
	startSuspensionTime := time.Date(parkTime.Year(), parkTime.Month(), parkTime.Day(), SuspensionStartTime, 0, 0, 0, time.UTC)
	if leaveTime.Before(startSuspensionTime) {
		totalPrice += CalculateParkingPriceWithTime(leaveTime.Sub(parkTime).Minutes(), PricePerHour)
		if leaveTime.Sub(parkTime).Minutes() > freeParkingTime() {
			totalPrice -= FreeParkingDiscount
			return totalPrice
		}
		return totalPrice
	}

	totalPrice += SuspensionFee
	if startSuspensionTime.Sub(parkTime).Minutes() > freeParkingTime() {
		totalPrice += CalculateParkingPriceWithTime(startSuspensionTime.Sub(parkTime).Minutes(), PricePerHour)
		totalPrice -= FreeParkingDiscount
	}

	endSuspensionTime := time.Date(leaveTime.Year(), leaveTime.Month(), leaveTime.Day(), SuspensionEndTime, 0, 0, 0, time.UTC)
	if leaveTime.After(endSuspensionTime) && parkTime.Day() != leaveTime.Day() {
		totalPrice += CalculateParkingPriceWithTime(leaveTime.Sub(endSuspensionTime).Minutes(), PricePerHour)
		return totalPrice
	}

	return totalPrice
}

func CalculateMultipleDayParkingFee2(parkTime, leaveTime time.Time) (totalPrice int) {
	startSuspensionTime := time.Date(parkTime.Year(), parkTime.Month(), parkTime.Day(), SuspensionStartTime, 0, 0, 0, time.UTC)
	if startSuspensionTime.Sub(parkTime).Minutes() > freeParkingTime() {
		totalPrice += CalculateParkingPriceWithTime(startSuspensionTime.Sub(parkTime).Minutes(), PricePerHour)
		totalPrice -= FreeParkingDiscount
	}

	dayDifference := int(leaveTime.Sub(parkTime).Hours() / 24)
	if dayDifference > 0 {
		totalPrice += SuspensionFee * dayDifference
		totalPrice += DailyParkingFeeBeforeSuspension * (dayDifference - 1)
	}

	endSuspensionTime := time.Date(leaveTime.Year(), leaveTime.Month(), leaveTime.Day(), SuspensionEndTime, 0, 0, 0, time.UTC)
	roundUpLeaveTime := leaveTime
	if leaveTime.Minute() != 0 || leaveTime.Second() != 0 {
		roundUpLeaveTime = leaveTime.Add(time.Hour).Truncate(time.Hour)
	}

	if roundUpLeaveTime.Hour() <= SuspensionStartTime && endSuspensionTime.Before(roundUpLeaveTime) {
		totalPrice += int(math.Abs(endSuspensionTime.Sub(roundUpLeaveTime).Hours())) * PricePerHour
	} else {
		totalPrice += DailyParkingFeeBeforeSuspension
	}

	if isSuspensionTime(leaveTime) {
		totalPrice += SuspensionFee
	}

	return totalPrice
}

func freeParkingTime() float64 {
	return time.Hour.Minutes() * FreeParkingHours
}

func CalculateParkingPriceWithTime(actualMinutes float64, pricePerHour int) int {
	return int(math.Ceil(actualMinutes/time.Hour.Minutes())) * pricePerHour
}

func isSuspensionTime(leaveTime time.Time) bool {
	return leaveTime.Hour() >= SuspensionStartTime || leaveTime.Hour() < SuspensionEndTime
}
