package processing

import (
	"context"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"
	"os"
)

func CreateEmbeddings(issues []string) [][]float32 {
	embeddings := make([][]float32, len(issues))
	for i := 0; i < len(issues); i++ {
		embeddings[i] = generateEmbedding(issues[i])
	}
	return embeddings
}

// Generates a vector embedding for a single input string using the Google Gemini Pro embedding model ("embedding-001")
func generateEmbedding(issue string) []float32 {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	em := client.EmbeddingModel("embedding-001")
	res, err := em.EmbedContent(ctx, genai.Text(issue))
	if err != nil {
		log.Fatal(err)
	}
	return res.Embedding.Values
}
