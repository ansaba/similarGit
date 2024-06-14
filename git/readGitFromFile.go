package git

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadIssues() []Issue {
	data, err := os.ReadFile("github_api_dump.json") // Read entire file
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil // Exit with an error
	}

	// Parse the JSON response
	var issues []Issue
	if err := json.Unmarshal(data, &issues); err != nil {
		fmt.Printf("Failed to parse issues: %v", err)
	}

	return issues

}
