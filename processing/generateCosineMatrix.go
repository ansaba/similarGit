package processing

import (
	"similarGit/git"
	"similarGit/ml"
)

func GenerateCosineMatrix(embeddings [][]float32, issues []git.Issue) map[string][]string {
	dupeIndex := make(map[string][]string)
	cosine_matrix := make([][]float32, len(embeddings))
	for i := range cosine_matrix {
		cosine_matrix[i] = make([]float32, len(embeddings))
	}

	for i := 0; i < len(cosine_matrix); i++ {
		for j := 0; j < len(cosine_matrix); j++ {
			cosine_matrix[i][j], _ = ml.CosineDistance(embeddings[i], embeddings[j])
			if cosine_matrix[i][j] > .8 && (i != j) {
				addToIndex(dupeIndex, i, j, issues)
			}
		}
	}
	return dupeIndex
}

func addToIndex(dupeIndex map[string][]string, i int, j int, issues []git.Issue) {
	var issue1 string = issues[i].HTMLURL
	var issue2 string = issues[j].HTMLURL
	val, exists := dupeIndex[issue1]
	if exists {
		dupeIndex[issue1] = append(val, issue2)
	} else {
		dupeIndex[issue1] = []string{issue2}
	}
}
