package main

import (
	"fmt"
	"github.com/google/go-github/v62/github"
	"regexp"
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

func printPrettySearchIndex(dupeIndex map[string][]string) {
	fmt.Println("\nSimilar issues")
	for key, values := range dupeIndex {
		// Extract issue number from the key
		keyNumber := ExtractIssueNumber(key)

		// Extract issue numbers from the values
		var valueNumbers []string
		for _, value := range values {
			num := ExtractIssueNumber(value)
			valueNumbers = append(valueNumbers, num)
		}
		// Print the results
		fmt.Printf(" %s: %v\n", keyNumber, valueNumbers)
	}
}

// ExtractIssueNumber extracts the issue number from a GitHub issue URL.
func ExtractIssueNumber(url string) string {
	// Regular expression to match GitHub issue URL and capture the issue number
	re := regexp.MustCompile(`github\.com/.*/issues/(\d+)`)
	matches := re.FindStringSubmatch(url)
	if len(matches) < 2 {
		return ""
	}
	return matches[1]
}
