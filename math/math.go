package math

import (
	"math"
)

func IntRatio(a, b int64) float64 {
	if a == 0 || b == 0 {
		return 0
	}

	return float64(a) / float64(b)
}

func IntRatioPercentage(a, b int64) float64 {
	return IntRatio(a, b) * 100
}

func RoundToNdp(f, noDP float64) float64 {
	exp := math.Pow(10, noDP)
	return math.Round(f*exp) / exp
}
