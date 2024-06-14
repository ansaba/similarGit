package git

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	// GitHub API endpoint for issues
	githubIssuesURL = "https://api.github.com/repos/golang/go/issues?per_page=3"
	//	githubIssuesURL = "https://api.github.com/repos/golang/go/issues?q=is:issue is:open project:golang/43"
)

func joinLabels(labels []Label) string {
	var labelNames []string
	for _, label := range labels {
		labelNames = append(labelNames, label.Name)
	}
	return strings.Join(labelNames, ", ")
}

func fetchIssues() []Issue {
	fmt.Printf("Making get call %s\n", githubIssuesURL)
	// Make the GET request to the GitHub API
	resp, err := http.Get(githubIssuesURL)
	if err != nil {
		fmt.Printf("Failed to fetch issues: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Failed to fetch issues: %s", resp.Status)
	}

	// Check if the response body is empty
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v", err)
	}

	if len(body) == 0 {
		fmt.Printf("Response body is empty")
	}

	// Save response body to a file for reuse
	err = saveToFile("github_api_response.json", body)
	if err != nil {
		fmt.Println("Error saving to file:", err)
		return nil
	}

	// Parse the JSON response
	var issues []Issue
	//if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
	if err := json.Unmarshal(body, &issues); err != nil {
		fmt.Println("Failed to parse issues: %v", err)
	}

	fmt.Printf("Parsing succeeded\n")

	return issues
}

func saveToFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
