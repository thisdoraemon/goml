package main

import (
	"fmt"
)

// Calculate the slope (beta1) of the linear regression model
func calculateSlope(x []float64, meanX float64, y []float64, meanY float64) float64 {
	numerator := 0.0
	denominator := 0.0

	for i := 0; i < len(x); i++ {
		numerator += (x[i] - meanX) * (y[i] - meanY)
		denominator += (x[i] - meanX) * (x[i] - meanX)
	}

	return numerator / denominator
}

// Calculate the mean of a slice of floats
func mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

func main() {
	// Sample data: house sizes (in square feet) and their corresponding prices (in thousands of dollars)
	sizes := []float64{1400, 1600, 1700, 1875, 1100, 1550, 2350, 2450, 1425, 1700}
	prices := []float64{245, 312, 279, 308, 199, 219, 405, 324, 319, 255}

	// Calculate the mean of sizes and prices
	meanSize := mean(sizes)
	meanPrice := mean(prices)

	// Calculate the slope (beta1) and intercept (beta0) of the linear regression model
	beta1 := calculateSlope(sizes, meanSize, prices, meanPrice)
	beta0 := meanPrice - (beta1 * meanSize)

	// Predict the price for a house with a size of 2000 square feet
	houseSize := 2000.0
	predictedPrice := beta0 + (beta1 * houseSize)

	// Print the results
	fmt.Printf("Intercept (beta0): %.2f\n", beta0)
	fmt.Printf("Slope (beta1): %.2f\n", beta1)
	fmt.Printf("Predicted Price for a %.2f sq. ft. house: $%.2f\n", houseSize, predictedPrice)
}
