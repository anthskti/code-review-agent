package models

type Finding struct {
	Severity    string `json:"severity"`
	IssueType   string `json:"issue_type"`
	Description string `json:"description"`
	Action      string `json:"action"`
	LineNumber  *int   `json:"line_number,omitempty"`
}
