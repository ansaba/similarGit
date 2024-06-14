package processing

import (
	"encoding/json"
	"fmt"
	"os"
	"similarGit/git"
)

func Preprocess(issues []git.Issue) []string {
	var processedIssues []string
	for _, issue := range issues {
		processedIssues = append(processedIssues, preprocessIssue(issue))
	}
	return processedIssues
}

func preprocessIssue(issue git.Issue) string {
	var lLMIssue = covertIssuetoLLMRelevantIssue(issue) // Marshal the struct into JSON

	jsonData, err := json.MarshalIndent(lLMIssue, "", "  ") // Indented for readability
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		os.Exit(1) // Exit with an error code
	}
	var jsonString = string(jsonData) // Print the JSON output
	//	fmt.Println("Preprocessed JSON : %s", jsonString)
	return jsonString
}

// TODO : Extract project prefix before
// Extract relevant nodes for LLM from Github Issues
func covertIssuetoLLMRelevantIssue(issue git.Issue) git.LLMRelevantIssue {
	var lLMIssue git.LLMRelevantIssue
	lLMIssue.Title = issue.Body
	lLMIssue.Title = issue.Title
	lLMLabels := make([]git.LLMRelevantLabel, len(issue.Labels)) // Preallocate the new array

	for i := 0; i < len(issue.Labels); i++ {
		lLMLabels[i].Name = issue.Labels[i].Name
	}
	return lLMIssue
}
