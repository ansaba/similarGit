package git

// Issue represents a GitHub issue
type Issue struct {
	URL           string  `json:"url"`
	RepositoryURL string  `json:"repository_url"`
	HTMLURL       string  `json:"html_url"`
	Number        int     `json:"number"`
	State         string  `json:"state"`
	Title         string  `json:"title"`
	Body          string  `json:"body"`
	User          User    `json:"user"`
	Labels        []Label `json:"labels"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
	ClosedAt      *string `json:"closed_at"`
	Assignee      *User   `json:"assignee"`
}

// User represents a GitHub user
type User struct {
	Login   string `json:"login"`
	ID      int    `json:"id"`
	HTMLURL string `json:"html_url"`
}

// Label represents a label on a GitHub issue
type Label struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Default bool   `json:"default"`
}
