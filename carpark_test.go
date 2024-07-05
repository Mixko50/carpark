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
		result := CalculateParkingFee(test.parkTime, test.leaveTime)
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
		result := CalculateParkingFee(test.parkTime, test.leaveTime)
		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}
