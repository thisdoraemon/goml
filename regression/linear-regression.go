package regression

func CalculateSlope(x []float64, meanX float64, y []float64, meanY float64) float64 {
	numerator := 0.0
	denominator := 0.0

	for i := 0; i < len(x); i++ {
		numerator += (x[i] - meanX) * (y[i] - meanY)
		denominator += (x[i] - meanX) * (x[i] - meanX)
	}

	return numerator / denominator
}

func Mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}
