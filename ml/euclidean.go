package ml

import (
	"fmt"
	"math"
)

func EuclideanDistance(p1, p2 []float32) (float64, error) {
	if len(p1) != len(p2) {
		return 0, fmt.Errorf("vectors must have the same dimensions")
	}

	var sumSquares float32
	for i := range p1 {
		diff := p1[i] - p2[i]
		sumSquares += diff * diff
	}

	return math.Sqrt(float64(sumSquares)), nil
}
