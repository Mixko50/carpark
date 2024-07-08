package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type ParkingTicket struct {
	EntryTime int64 `json:"entry_time"`
	jwt.RegisteredClaims
}

func GenerateParkingTicket() (string, error) {
	secretKey := os.Getenv("PARKING_SECRET")

	if time.Now().Hour() >= SuspensionStartTime || time.Now().Hour() < SuspensionEndTime {
		return "", fmt.Errorf("cannot generate parking ticket after %dPM and before %dAM", SuspensionStartTime, SuspensionEndTime)
	}

	claims := ParkingTicket{
		//EntryTime: time.Now().Unix(),
		time.Date(2024, time.July, 6, 20, 0, 0, 0, time.UTC).Unix(),
		jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	// For testing
	writeTokenToFileFile(tokenString)

	return tokenString, nil
}

func writeTokenToFileFile(token string) {
	f, err := os.Create("./token.txt")
	if err != nil {
		fmt.Println("Cannot create file")
		panic(err)
	}

	n3, err := f.WriteString(token)
	if err != nil {
		fmt.Println("Cannot open file")
		panic(err)
	}

	fmt.Printf("wrote %d bytes\n", n3)
	f.Sync()
}
