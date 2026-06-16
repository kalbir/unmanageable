package tests

import (
	"testing"

	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/scoring"
	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/types"
)

func TestScorePoint_AllPassed(t *testing.T) {
	pt := &types.StandardPoint{
		Number: 9,
		Checks: []types.CheckResult{
			{ID: "c1", Score: 1.0},
			{ID: "c2", Score: 1.0},
		},
	}
	scoring.ScorePoint(pt)

	if pt.Score != 100 {
		t.Errorf("expected score 100, got %.1f", pt.Score)
	}
	if pt.Grade != types.GradeA {
		t.Errorf("expected grade A, got %s", pt.Grade)
	}
}

func TestScorePoint_AllFailed(t *testing.T) {
	pt := &types.StandardPoint{
		Number: 5,
		Checks: []types.CheckResult{
			{ID: "c1", Score: 0},
			{ID: "c2", Score: 0},
		},
	}
	scoring.ScorePoint(pt)

	if pt.Score != 0 {
		t.Errorf("expected score 0, got %.1f", pt.Score)
	}
	if pt.Grade != types.GradeF {
		t.Errorf("expected grade F, got %s", pt.Grade)
	}
}

func TestScorePoint_NoChecks(t *testing.T) {
	pt := &types.StandardPoint{Number: 1}
	scoring.ScorePoint(pt)

	if pt.Grade != types.GradeF {
		t.Errorf("expected grade F for no checks, got %s", pt.Grade)
	}
}

func TestScorePoint_SkipsErrors(t *testing.T) {
	pt := &types.StandardPoint{
		Number: 9,
		Checks: []types.CheckResult{
			{ID: "c1", Score: 1.0},
			{ID: "c2", Score: 0, Error: "connection refused"},
		},
	}
	scoring.ScorePoint(pt)

	if pt.Score != 100 {
		t.Errorf("expected score 100 (errors skipped), got %.1f", pt.Score)
	}
}

func TestScoreAssessment_AggregatesPoints(t *testing.T) {
	result := &types.AssessmentResult{
		Points: []types.StandardPoint{
			{
				Number: 9,
				Checks: []types.CheckResult{{Score: 1.0}, {Score: 1.0}},
			},
			{
				Number: 14,
				Checks: []types.CheckResult{{Score: 0.5}, {Score: 0.5}},
			},
		},
	}
	scoring.ScoreAssessment(result)

	if result.OverallScore < 60 {
		t.Errorf("unexpected overall score: %.1f", result.OverallScore)
	}
	if result.OverallGrade == "" {
		t.Error("expected an overall grade to be set")
	}
}

func TestScoreAssessment_EmptyPoints(t *testing.T) {
	result := &types.AssessmentResult{}
	scoring.ScoreAssessment(result)

	if result.OverallGrade != types.GradeF {
		t.Errorf("expected grade F for empty result, got %s", result.OverallGrade)
	}
}
