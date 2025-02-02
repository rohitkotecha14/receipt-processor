package services

import (
	"fmt"
	"math"
	"receipt-processor/models"
	"strconv"
	"strings"
	"time"
	"unicode"
	_ "unicode"
)

func CalculatePoints(receipt models.Receipt) (int, error) {
	totalPoints := 0

	totalPoints += calculateRetailerPoints(receipt.Retailer)

	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid total format")
	}

	totalPoints += calculateRoundDollarPoints(total)
	totalPoints += calculateMultipleQuarterPoints(total)
	totalPoints += calculateItemPairPoints(len(receipt.Items))

	itemPoints, err := calculateItemPoints(receipt.Items)
	if err != nil {
		return 0, err
	}
	totalPoints += itemPoints

	datePoints, err := calculatePurchaseDatePoints(receipt.PurchaseDate)
	if err != nil {
		return 0, err
	}
	totalPoints += datePoints

	timePoints, err := calculatePurchaseTimePoints(receipt.PurchaseTime)
	if err != nil {
		return 0, err
	}
	totalPoints += timePoints

	return totalPoints, nil
}

func calculateRetailerPoints(retailer string) int {
	points := 0
	for _, ch := range retailer {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			points++
		}
	}
	return points
}

func calculateRoundDollarPoints(total float64) int {
	if math.Mod(total, 1.0) == 0 {
		return 50
	}
	return 0
}

func calculateMultipleQuarterPoints(total float64) int {
	if math.Mod(total, 0.25) == 0 {
		return 25
	}
	return 0
}

func calculateItemPairPoints(itemCount int) int {
	return (itemCount / 2) * 5
}

func calculateItemPoints(items []models.Item) (int, error) {
	points := 0
	for _, item := range items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid item price format")
			}
			points += int(math.Ceil(price * 0.2))
		}
	}
	return points, nil
}

func calculatePurchaseDatePoints(purchaseDate string) (int, error) {
	parsedDate, err := time.Parse("2006-01-02", purchaseDate)
	if err != nil {
		return 0, fmt.Errorf("invalid purchaseDate format")
	}
	if parsedDate.Day()%2 == 1 {
		return 6, nil
	}
	return 0, nil
}

func calculatePurchaseTimePoints(purchaseTime string) (int, error) {
	parsedTime, err := time.Parse("15:04", purchaseTime)
	if err != nil {
		return 0, fmt.Errorf("invalid purchaseTime format")
	}
	hour := parsedTime.Hour()
	if hour == 14 || hour == 15 {
		return 10, nil
	}
	return 0, nil
}
