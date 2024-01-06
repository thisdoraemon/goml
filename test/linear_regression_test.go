package test

import (
	"fmt"
	"math"
	"testing"

	"github.com/thisdoraemon/goml/regression"
)

func almostEqual(a, b float64) bool {
	epsilon := 0.0001
	return math.Abs(a-b) < epsilon
}

func TestCalculateSlope(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 5, 4, 5}
	meanX := regression.Mean(x)
	meanY := regression.Mean(y)

	expectedResult := 0.6

	result := regression.CalculateSlope(x, meanX, y, meanY)

	if !almostEqual(result, expectedResult) {
		t.Errorf("Expected slope: %v, got: %v", expectedResult, result)
	}
}

func TestMean(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}

	expectedResult := 3.0

	result := regression.Mean(data)

	if result != expectedResult {
		t.Errorf("Expected mean: %v, got: %v", expectedResult, result)
	}
}

func TestPredictedPrice(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 5, 4, 5}

	meanX := regression.Mean(x)
	meanY := regression.Mean(y)

	slope := regression.CalculateSlope(x, meanX, y, meanY)

	fmt.Println(meanX)
	fmt.Println(meanY)

	predictedPrice := slope*(6-meanX) + meanY
	tolerance := 0.000001

	fmt.Println(predictedPrice)
	expectedPrice := 4.5

	if math.Abs(predictedPrice-expectedPrice) > tolerance {
		t.Errorf("predicted price for x = 6 is incorrect. Got %f, expected %f", predictedPrice, 4.5)
	}
}
