package main

import (
	"fmt"
	"github.com/google/go-github/v62/github"
)

func printSDKIssues(issues []*github.Issue) {
	for _, issue := range issues {
		fmt.Println("--------------------")
		fmt.Println("Number:", *issue.Number)
		fmt.Println("Title:", *issue.Title)
		fmt.Println("URL:", *issue.HTMLURL)
	}
}

func printMatrix(matrix [][]float32) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Printf("%4d", matrix[i][j])
		}
		fmt.Println()
	}
}

func printMatrix64(matrix [][]float64) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Printf("%4d", matrix[i][j])
		}
		fmt.Println()
	}
}

func printCosIndex(dupeIndex map[string][]string) {
	fmt.Println("\nSimilar issues")
	for key, values := range dupeIndex {
		fmt.Printf(" %s: %v\n", key, values)
	}
}
