package git

// LLMRelevantIssue represents a GitHub issue
type LLMRelevantIssue struct {
	Title  string             `json:"title"`
	Body   string             `json:"body"`
	Labels []LLMRelevantLabel `json:"labels"`
}

// LLMRelevantLabel represents a label on a GitHub issue
type LLMRelevantLabel struct {
	Name string `json:"name"`
}
