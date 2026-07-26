package models

type ReviewState struct {
	CodeSnippet string `json:"code_snippet"`
	SourceType  string `json:"source_type"`
	SourceLabel string `json:"source_label"`

	Context             *Context  `json:"context,omitempty"`
	SecurityFindings    []Finding `json:"security_findings"`
	PerformanceFindings []Finding `json:"performance_findings"`

	FinalReport string `json:"final_report"`
}
