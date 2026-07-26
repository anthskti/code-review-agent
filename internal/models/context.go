package models

type Context struct {
	Summary            string   `json:"summary"`
	DataClassification string   `json:"data_classification"`
	RiskProfile        string   `json:"risk_profile"`
	Dependencies       []string `json:"dependencies"`
	ApplicationType    string   `json:"application_type"`
	TrustBoundary      string   `json:"trust_boundary"`
}
