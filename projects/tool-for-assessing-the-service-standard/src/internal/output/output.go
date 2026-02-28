package output

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/types"
)

// Format is the output format identifier.
type Format string

const (
	FormatJSON Format = "json"
	FormatText Format = "text"
)

// WriteJSON serialises the result as indented JSON to w.
func WriteJSON(w io.Writer, result *types.AssessmentResult) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(result)
}

// WriteText writes a human-readable assessment report to w.
func WriteText(w io.Writer, result *types.AssessmentResult) error {
	fmt.Fprintf(w, "GDS Service Standard Assessment\n")
	fmt.Fprintf(w, "%s\n\n", strings.Repeat("=", 50))
	fmt.Fprintf(w, "Service:      %s\n", result.ServiceName)
	fmt.Fprintf(w, "URL:          %s\n", result.ServiceURL)
	fmt.Fprintf(w, "Assessed at:  %s\n", result.AssessedAt.Format("2006-01-02 15:04:05 UTC"))
	fmt.Fprintf(w, "Overall:      %s (%.1f%%)\n\n", result.OverallGrade, result.OverallScore)

	if len(result.Errors) > 0 {
		fmt.Fprintf(w, "Assessment errors:\n")
		for _, e := range result.Errors {
			fmt.Fprintf(w, "  ! %s\n", e)
		}
		fmt.Fprintln(w)
	}

	for _, pt := range result.Points {
		gradeStr := string(pt.Grade)
		if len(pt.Checks) == 0 {
			gradeStr = "N/A"
		}
		fmt.Fprintf(w, "Standard %2d: %s\n", pt.Number, pt.Title)
		fmt.Fprintf(w, "  Grade: %s (%.1f%%)\n", gradeStr, pt.Score)

		for _, c := range pt.Checks {
			status := "✓"
			if !c.Passed {
				status = "✗"
			}
			if c.Error != "" {
				status = "?"
			}
			fmt.Fprintf(w, "  %s %s\n", status, c.Name)
			if c.Evidence != "" {
				fmt.Fprintf(w, "      Evidence:    %s\n", c.Evidence)
			}
			if c.Error != "" {
				fmt.Fprintf(w, "      Error:       %s\n", c.Error)
			}
			if !c.Passed && c.Remediation != "" {
				fmt.Fprintf(w, "      Remediation: %s\n", c.Remediation)
			}
		}

		if len(pt.Checks) == 0 {
			fmt.Fprintf(w, "  (no automated checks available for this standard point in v1)\n")
		}
		fmt.Fprintln(w)
	}

	return nil
}
