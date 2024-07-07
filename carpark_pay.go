package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

func GetParkingFee(token string) (int, error) {
	secretKey := os.Getenv("PARKING_SECRET")

	// Parse the token
	var claims ParkingTicket
	decoded, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC and HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Fatalf("Failed to parse token: %v", err)
	}
	if decoded.Valid {
	} else {
		log.Fatal("Token is invalid")
	}

	fmt.Println(claims.EntryTime)
	parkTime := time.Unix(claims.EntryTime, 0)
	parkTime = time.Date(parkTime.Year(), parkTime.Month(), parkTime.Day(), parkTime.Hour()-7, parkTime.Minute(), parkTime.Second(), 0, time.UTC)
	currentTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC)

	// Calculate the parking fee
	result := CalculateParkingFee(parkTime, currentTime)
	return result, nil
}
