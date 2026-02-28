package types

import "time"

// Grade represents a letter grade for an assessment result.
type Grade string

const (
	GradeA Grade = "A"
	GradeB Grade = "B"
	GradeC Grade = "C"
	GradeD Grade = "D"
	GradeF Grade = "F"
)

// CheckResult holds the outcome of a single check.
type CheckResult struct {
	ID          string
	Name        string
	Description string
	Passed      bool
	Score       float64 // 0.0 to 1.0
	Evidence    string
	Remediation string
	Automated   bool
	Error       string
}

// StandardPoint maps to one of the 14 GDS service standard points.
type StandardPoint struct {
	Number      int
	Title       string
	Description string
	Checks      []CheckResult
	Grade       Grade
	Score       float64
}

// AssessmentResult is the top-level result for a single service.
type AssessmentResult struct {
	ServiceName   string
	ServiceURL    string
	AssessedAt    time.Time
	OverallGrade  Grade
	OverallScore  float64
	Points        []StandardPoint
	AutomatedOnly bool
	Errors        []string
}

// GradeFromScore converts a 0–100 score to a letter grade.
func GradeFromScore(score float64) Grade {
	switch {
	case score >= 90:
		return GradeA
	case score >= 75:
		return GradeB
	case score >= 60:
		return GradeC
	case score >= 40:
		return GradeD
	default:
		return GradeF
	}
}
