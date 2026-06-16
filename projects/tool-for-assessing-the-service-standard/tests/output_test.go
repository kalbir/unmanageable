package tests

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/output"
	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/types"
)

func makeTestResult() *types.AssessmentResult {
	return &types.AssessmentResult{
		ServiceName:  "Test Service",
		ServiceURL:   "https://example.gov.uk",
		AssessedAt:   time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC),
		OverallGrade: types.GradeB,
		OverallScore: 80.0,
		Points: []types.StandardPoint{
			{
				Number: 9,
				Title:  "Create a secure service which protects users' privacy",
				Grade:  types.GradeA,
				Score:  90.0,
				Checks: []types.CheckResult{
					{ID: "web-https", Name: "HTTPS enforced", Passed: true, Score: 1.0, Evidence: "TLS 1.3"},
				},
			},
		},
	}
}

func TestWriteJSON_ValidJSON(t *testing.T) {
	result := makeTestResult()
	var buf bytes.Buffer
	if err := output.WriteJSON(&buf, result); err != nil {
		t.Fatalf("WriteJSON returned error: %v", err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &parsed); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}

	if parsed["ServiceName"] != "Test Service" {
		t.Errorf("unexpected service name in JSON: %v", parsed["ServiceName"])
	}
}

func TestWriteText_ContainsServiceName(t *testing.T) {
	result := makeTestResult()
	var buf bytes.Buffer
	if err := output.WriteText(&buf, result); err != nil {
		t.Fatalf("WriteText returned error: %v", err)
	}

	text := buf.String()
	if !strings.Contains(text, "Test Service") {
		t.Error("text output does not contain service name")
	}
	if !strings.Contains(text, "https://example.gov.uk") {
		t.Error("text output does not contain service URL")
	}
	if !strings.Contains(text, "B") {
		t.Error("text output does not contain overall grade")
	}
}

func TestWriteText_ShowsRemediation(t *testing.T) {
	result := makeTestResult()
	result.Points[0].Checks = []types.CheckResult{
		{
			ID:          "web-https",
			Name:        "HTTPS enforced",
			Passed:      false,
			Score:       0,
			Evidence:    "URL does not use HTTPS",
			Remediation: "Serve the service over HTTPS.",
		},
	}

	var buf bytes.Buffer
	if err := output.WriteText(&buf, result); err != nil {
		t.Fatalf("WriteText returned error: %v", err)
	}

	text := buf.String()
	if !strings.Contains(text, "Remediation") {
		t.Error("text output should contain remediation for failed check")
	}
}

func TestWriteText_NoAutomatedChecks(t *testing.T) {
	result := makeTestResult()
	result.Points[0].Checks = nil

	var buf bytes.Buffer
	if err := output.WriteText(&buf, result); err != nil {
		t.Fatalf("WriteText returned error: %v", err)
	}

	text := buf.String()
	if !strings.Contains(text, "no automated checks") {
		t.Error("expected message about no automated checks")
	}
}
