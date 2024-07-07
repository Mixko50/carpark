package main

import (
	"math"
	"time"
)

const (
	pricePerHour        = 100
	freeParkingHours    = 2
	suspensionStartTime = 22
	suspensionEndTime   = 10
	suspensionFee       = 1000
	dailyParkingFee     = 2200
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
		return CalculateMultipleDayParkingFee(parkTime, leaveTime)
	}

	return CalculateSingleDayParkingFee(parkTime, leaveTime)
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

func CalculateMultipleDayParkingFee(parkTime, leaveTime time.Time) int {
	startSuspensionTime := time.Date(parkTime.Year(), parkTime.Month(), parkTime.Day(), suspensionStartTime, 0, 0, 0, time.UTC)
	difference := leaveTime.Sub(parkTime)
	// Cut the free parking time
	actualMinutes := difference.Minutes() - freeParkingTime()
	totalPrice := 0
	// Find day difference
	dayDifference := int(leaveTime.Sub(parkTime).Hours() / 24)

	// Calculate price for each day
	// 2200 is the price for 24 hours without free parking time
	totalPrice += dailyParkingFee * dayDifference

	// If the park time is not 21:00, then we need to deduct 200 baht for first 2 hours of parking fee
	if dayDifference > 1 && parkTime.Hour() != 21 {
		totalPrice -= 200
	}

	// If the leave time is 21:00, then we need to deduct 100 baht for only first an hour of parking fee
	if parkTime.Hour() == 21 {
		totalPrice -= 100
	}

	// Set the park time to be the day as leave time
	tempParkTime := parkTime.Add(time.Hour * 24 * time.Duration(dayDifference))

	// Calculate the actual minutes for the last day
	actualMinutes = math.Abs(leaveTime.Sub(tempParkTime).Minutes())

	if isSuspensionTime(leaveTime) {
		totalPrice += suspensionFee

		// Cut the suspension time
		actualMinutes -= leaveTime.Sub(time.Date(leaveTime.Year(), leaveTime.Month(), leaveTime.Day(), suspensionStartTime, 0, 0, 0, time.UTC)).Minutes()
	}
	if startSuspensionTime.Sub(parkTime).Minutes() > freeParkingTime() {
		tempActualMinutes := startSuspensionTime.Sub(parkTime).Minutes()
		totalPrice += CalculateParkingPriceWithTime(tempActualMinutes, pricePerHour)
	} else {
		totalPrice += CalculateParkingPriceWithTime(actualMinutes, pricePerHour)
	}
	return totalPrice
}
