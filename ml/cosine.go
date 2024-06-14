package ml

import (
	"fmt"
	"math"
)

func dotProduct(v1, v2 []float32) (float32, error) {
	if len(v1) != len(v2) {
		return 0, fmt.Errorf("vectors must have the same length")
	}

	var res float32
	for i := range v1 {
		res += v1[i] * v2[i]
	}
	return res, nil
}

func magnitude(v []float32) float32 {
	var sumOfSquares float32
	for _, val := range v {
		sumOfSquares += val * val
	}
	return float32(math.Sqrt(float64(sumOfSquares)))
}

func CosineDistance(vector1 []float32, vector2 []float32) (float32, error) {
	// Calculate the dot product of the two vectors
	dotProd, err := dotProduct(vector1, vector2)
	if err != nil {
		return 0, err
	}

	mag1 := magnitude(vector1)
	mag2 := magnitude(vector2)
	if mag1 == 0 || mag2 == 0 {
		return 0, fmt.Errorf("magnitude of a vector cannot be zero")
	}

	// Calculate the cosine similarity between the two vectors
	cosineSimilarity := dotProd / (mag1 * mag2)

	return cosineSimilarity, nil
}
