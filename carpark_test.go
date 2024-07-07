package main

import (
	"testing"
	"time"
)

func TestParkLessThanTwoHours(t *testing.T) {
	type testCase struct {
		name      string
		parkTime  time.Time
		leaveTime time.Time
		expected  int
	}

	tests := []testCase{
		{
			// 1 hour
			"1 hour",
			time.Date(2019, time.January, 1, 10, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 11, 0, 0, 0, time.UTC),
			0,
		},
		{
			// 1 hour 59 minutes
			"1 hour 59 minutes",
			time.Date(2019, time.January, 1, 10, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 11, 59, 0, 0, time.UTC),
			0,
		},
		{
			// 2 hours
			"2 hours",
			time.Date(2019, time.January, 1, 10, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 12, 0, 0, 0, time.UTC),
			0,
		},
		{
			"1 hour 3 minutes 2 seconds",
			time.Date(2019, time.January, 1, 10, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 1, 3, 2, 0, time.UTC),
			0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := CalculateParkingFee(test.parkTime, test.leaveTime)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}

func TestParkMoreThanTwoHoursAndBeforeTenPm(t *testing.T) {
	type testCase struct {
		name      string
		parkTime  time.Time
		leaveTime time.Time
		expected  int
	}
	tests := []testCase{
		{
			// 2 hours
			"2 hours",
			time.Date(2019, time.January, 1, 20, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 22, 0, 0, 0, time.UTC),
			0,
		},
		{
			// 3 hours
			"3 hours",
			time.Date(2019, time.January, 1, 10, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 13, 0, 0, 0, time.UTC),
			100,
		},
		{
			// 6 hours
			"6 hours",
			time.Date(2019, time.January, 1, 12, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 18, 0, 0, 0, time.UTC),
			400,
		},
		{
			// 4 hours 1 minute
			"4 hours 1 minute",
			time.Date(2019, time.January, 1, 11, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 15, 0, 1, 0, time.UTC),
			300,
		},
		{
			// 5 hours 25 minutes 52 seconds
			"5 hours 25 minutes 52 seconds",
			time.Date(2019, time.January, 1, 12, 10, 4, 0, time.UTC),
			time.Date(2019, time.January, 1, 17, 35, 56, 0, time.UTC),
			400,
		},
		{
			// 20 hours 23 minutes 2 seconds
			"20 hours 23 minutes 2 seconds",
			time.Date(2019, time.January, 1, 11, 5, 4, 0, time.UTC),
			time.Date(2019, time.January, 1, 21, 38, 6, 0, time.UTC),
			900,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := CalculateParkingFee(test.parkTime, test.leaveTime)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}

func TestParkMoreThanTwoHoursAndLeaveAfterTenPmButBeforeTenPm(t *testing.T) {
	type testCase struct {
		name      string
		parkTime  time.Time
		leaveTime time.Time
		expected  int
	}
	tests := []testCase{
		{
			// 2 hours 1 minute
			"2 hours 1 minute",
			time.Date(2019, time.January, 1, 20, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 22, 1, 0, 0, time.UTC),
			1000,
		},
		{
			// 3 hours
			"3 hours",
			time.Date(2019, time.January, 1, 20, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 1, 23, 0, 0, 0, time.UTC),
			1000,
		},
		{
			// 6 hours, day 1 from 20:00 to 23:59, day 2 from 00:00 to 02:00
			"6 hours",
			time.Date(2019, time.January, 1, 20, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 2, 2, 0, 0, 0, time.UTC),
			1000,
		},
		{
			// 8 hours, day 1 from 19:00 to 23:59, day 2 from 00:00 to 03:00
			"8 hours",
			time.Date(2019, time.January, 1, 19, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 2, 3, 0, 0, 0, time.UTC),
			1100,
		},
		{
			// 19 hours, day 1 from 19:00 to 23:59, day 2 from 00:00 to 10:00
			"19 hours",
			time.Date(2019, time.January, 1, 13, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 2, 8, 0, 0, 0, time.UTC),
			1700,
		},
		{
			// 23 hours, day 1 from 10:00 to 23:59, day 2 from 00:00 to 9:00
			"23 hours",
			time.Date(2019, time.January, 1, 10, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 2, 9, 0, 0, 0, time.UTC),
			2000,
		},
		{
			"23 hours, 32 minutes, 12 seconds",
			time.Date(2019, time.January, 1, 15, 7, 28, 0, time.UTC),
			time.Date(2019, time.January, 2, 14, 39, 40, 0, time.UTC),
			2000,

			// day 1: 15:07 - 17:07 -> Free
			// day 1: 17:08 - 21:59 -> 500
			// day 1: 22:00 - 09:59 -> 1000
			// day 2: 10:00 - 14:39 -> 500
			// 500 + 1000 + 500 = 2000
		},
		{
			"7 hours, 30 minutes, 12 seconds",
			time.Date(2019, time.January, 1, 21, 30, 28, 0, time.UTC),
			time.Date(2019, time.January, 2, 5, 0, 40, 0, time.UTC),
			1000,

			// day 1: 21:30 - 22:00 -> Free
			// day 1: 22:00 - 05:00 -> 1000
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := CalculateParkingFee(test.parkTime, test.leaveTime)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}

func TestParkMoreThan24Hours(t *testing.T) {
	type testCase struct {
		name      string
		parkTime  time.Time
		leaveTime time.Time
		expected  int
	}
	tests := []testCase{
		{
			"25 hours",
			time.Date(2019, time.January, 3, 21, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 4, 22, 0, 0, 0, time.UTC),
			3200,

			// day 1: 21 - 21:59 -> Free
			// day 2: 22 - 10 -> 1000
			// day 2: 10 - 21:59 -> 1200
			// day 3: 22 -> 1000
			// 1000 + 1200 + 1000 = 3200
		},
		{
			"49 hours",
			time.Date(2019, time.January, 1, 20, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 3, 21, 0, 0, 0, time.UTC),
			4300,

			// day 1: 20 - 22 -> Free
			// day 2: 22 - 10 -> 1000
			// day 2: 10 - 22 -> 1200
			// day 3: 22 - 10 -> 1000
			// day 3: 10 - 21 -> 1100
		},
		{
			"50 hours",
			time.Date(2019, time.January, 1, 20, 0, 0, 0, time.UTC),
			time.Date(2019, time.January, 3, 22, 0, 0, 0, time.UTC),
			5400,

			// day 1: 20 - 22 -> Free
			// day 2: 22 - 9:59 -> 1000
			// day 2: 10 - 21:59 -> 1200
			// day 3: 22 - 9:59 -> 1000
			// day 3: 10 - 21:59 -> 1200
			// day 3: 22 -> 1000
			// 1000 + 1200 + 1000 + 1200 + 1000 = 5400
		},
		{
			"98 hours",
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
		{
			"18 days, 12 hours, 19 minutes, 8 seconds",
			time.Date(2019, time.January, 1, 13, 45, 34, 0, time.UTC),
			time.Date(2019, time.January, 20, 2, 4, 42, 0, time.UTC),
			41300,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := CalculateParkingFee(test.parkTime, test.leaveTime)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}
