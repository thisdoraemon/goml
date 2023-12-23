package test

import (
	"github.com/thisdoraemon/goml/regression"
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	const epsilon = 1e-2
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
	sizes := []float64{1400, 1600, 1700, 1875, 1100, 1550, 2350, 2450, 1425, 1700}
	prices := []float64{245, 312, 279, 308, 199, 219, 405, 324, 319, 255}
	houseSize := 2000.0

	expectedIntercept := 67.62
	expectedSlope := 0.13
	expectedPredictedPrice := 309.62

	meanSize := regression.Mean(sizes)
	meanPrice := regression.Mean(prices)
	beta1 := regression.CalculateSlope(sizes, meanSize, prices, meanPrice)
	beta0 := meanPrice - (beta1 * meanSize)

	if !almostEqual(beta0, expectedIntercept) {
		t.Errorf("Expected intercept: %v, got: %v", expectedIntercept, beta0)
	}
	if !almostEqual(beta1, expectedSlope) {
		t.Errorf("Expected slope: %v, got: %v", expectedSlope, beta1)
	}

	predictedPrice := beta0 + (beta1 * houseSize)

	if !almostEqual(predictedPrice, expectedPredictedPrice) {
		t.Errorf("Expected predicted price: %v, got: %v", expectedPredictedPrice, predictedPrice)
	}
}
