package main

import (
	"similarGit/git"
	"similarGit/processing"
)

func main() {
	//Read the Github issues from JSON file
	issues := git.ReadIssues()

	// Process the data for creating embedding vectors
	processedIssues := processing.Preprocess(issues)

	//Create the Embeddings
	embeddings := processing.CreateEmbeddings(processedIssues)

	// Generate cosine Matrix and return similar issues search index
	duplicateIndex := processing.GenerateCosineMatrix(embeddings, issues)

	//Print the Similar issues index
	printPrettySearchIndex(duplicateIndex)
}
