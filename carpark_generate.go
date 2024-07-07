package main

import (
	"errors"
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
	if os.Getenv("PARKING_SECRET") == "" {
		return "", errors.New("PARKING_SECRET is not set")
	}

	secretKey := os.Getenv("PARKING_SECRET")

	claims := ParkingTicket{
		EntryTime: time.Now().Unix(),
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv(secretKey)))
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
