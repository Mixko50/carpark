package main

import (
	"testing"
	"time"
)

func TestParkLessThanTwoHours(t *testing.T) {
	type testCase struct {
		parkTime  time.Time
		leaveTime time.Time
		expected  int
	}

	tests := []testCase{
		{
			// 1 hour
			time.Date(2019, time.January, 1, 10, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 11, 0, 0, 0, time.UTC),
			0,
		},
		{
			// 1 hour 59 minutes
			time.Date(2019, time.January, 1, 10, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 11, 59, 0, 0, time.UTC),
			0,
		},
		{
			// 2 hours
			time.Date(2019, time.January, 1, 10, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 12, 0, 0, 0, time.UTC),
			0,
		},
	}

	for _, test := range tests {
		result, _ := CalculateParkingFee(test.parkTime, test.leaveTime)
		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}

func TestParkMoreThanTwoHours(t *testing.T) {
	type testCase struct {
		parkTime  time.Time
		leaveTime time.Time
		expected  int
	}
	tests := []testCase{
		{
			// 3 hours
			time.Date(2019, time.January, 1, 10, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 13, 0, 0, 0, time.UTC),
			100,
		},
		{
			// 6 hours
			time.Date(2019, time.January, 1, 12, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 18, 0, 0, 0, time.UTC),
			400,
		},
		{
			// 4 hours 1 minute
			time.Date(2019, time.January, 1, 11, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 15, 1, 0, 0, time.UTC),
			300,
		},
	}

	for _, test := range tests {
		result, _ := CalculateParkingFee(test.parkTime, test.leaveTime)
		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}

func TestParkMoreThanTwoHoursAndLeaveAfterTenPmButBeforeTenPm(t *testing.T) {
	type testCase struct {
		parkTime  time.Time
		leaveTime time.Time
		expected  int
	}
	tests := []testCase{
		{
			// 2 hours
			time.Date(2019, time.January, 1, 20, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 22, 0, 0, 0, time.UTC),
			0,
		},
		{
			// 2 hours 1 minute
			time.Date(2019, time.January, 1, 20, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 22, 1, 0, 0, time.UTC),
			1000,
		},
		{
			// 3 hours
			time.Date(2019, time.January, 1, 20, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 23, 0, 0, 0, time.UTC),
			1000,
		},
		{
			// 6 hours, day 1 from 20:00 to 23:59, day 2 from 00:00 to 02:00
			time.Date(2019, time.January, 1, 20, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 2, 2, 0, 0, 0, time.UTC),
			1000,
		},
		{
			// 8 hours, day 1 from 19:00 to 23:59, day 2 from 00:00 to 03:00
			time.Date(2019, time.January, 1, 19, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 2, 3, 0, 0, 0, time.UTC),
			1100,
		},
		{
			// 19 hours, day 1 from 19:00 to 23:59, day 2 from 00:00 to 10:00
			time.Date(2019, time.January, 1, 13, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 2, 8, 0, 0, 0, time.UTC),
			1700,
		},
		{
			// 24 hours, day 1 from 10:00 to 23:59, day 2 from 00:00 to 10:00
			time.Date(2019, time.January, 1, 10, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 2, 10, 0, 0, 0, time.UTC),
			2000,
		},
	}

	for _, test := range tests {
		result, _ := CalculateParkingFee(test.parkTime, test.leaveTime)
		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}

func TestParkMoreThan24Hours(t *testing.T) {
	type testCase struct {
		parkTime  time.Time
		leaveTime time.Time
		expected  int
	}
	tests := []testCase{
		{
			time.Date(2019, time.January, 1, 20, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 3, 21, 0, 0, 0, time.UTC),
			4300,
		},
		{
			time.Date(2019, time.January, 1, 20, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 3, 22, 0, 0, 0, time.UTC),
			5400,
		},
		{
			time.Date(2019, time.January, 1, 20, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 5, 23, 0, 0, 0, time.UTC),
			9800,

			// day 1: 20 - 22 -> Free
			// day 2: 22 - 10 -> 1000
			// day 2: 10 - 22 -> 1200
			// day 3: 22 - 10 -> 1000
			// day 3: 10 - 22 -> 1200
			// day 4: 22 - 10 -> 1000
			// day 4: 10 - 22 -> 1200
			// day 5: 22 - 10 -> 1000
			// day 5: 10 - 22 -> 1200
			// day 5: 22 - 23 -> 1000
			// 1000 + 1200 + 1000 + 1200 + 1000 + 1200 + 1000 + 1200 + 1000 = 9800
		},
	}

	for _, test := range tests {
		result, _ := CalculateParkingFee(test.parkTime, test.leaveTime)
		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}
