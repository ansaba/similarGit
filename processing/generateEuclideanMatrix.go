package processing

import (
	"similarGit/git"
	"similarGit/ml"
)

func generateEuclideanMatrix(embeddings [][]float32, issues []git.Issue) {
	euclidean_matrix := make([][]float64, len(embeddings))
	for i := range euclidean_matrix {
		euclidean_matrix[i] = make([]float64, len(embeddings))
	}

	for i := 0; i < len(euclidean_matrix); i++ {
		for j := 0; j < len(euclidean_matrix); j++ {
			euclidean_matrix[i][j], _ = ml.EuclideanDistance(embeddings[i], embeddings[j])
			if euclidean_matrix[i][j] < 1 {
				//	addToIndex(i, j, issues)
			}
			//	fmt.Printf("Euclidean matrix[%d][%d]: %f\n", i, j, euclidean_matrix[i][j])
		}
	}
	//	printMatrix64(euclidean_matrix)
}
