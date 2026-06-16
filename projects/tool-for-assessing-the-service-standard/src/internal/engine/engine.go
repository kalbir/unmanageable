package engine

import (
	"fmt"
	"time"

	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/checker"
	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/scoring"
	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/standard"
	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/types"
)

// Options configures an assessment run.
type Options struct {
	AutomatedOnly bool
	Verbose       bool
}

// Engine orchestrates all checkers and builds the final AssessmentResult.
type Engine struct {
	checkers []checker.Checker
}

// New returns an Engine with the default set of checkers.
func New() *Engine {
	return &Engine{
		checkers: []checker.Checker{
			checker.NewWebChecker(),
			checker.NewAccessibilityChecker(),
			checker.NewPerformanceChecker(),
			checker.NewPrivacyChecker(),
		},
	}
}

// Assess runs all checks against serviceURL and returns a populated AssessmentResult.
func (e *Engine) Assess(serviceName, serviceURL string, opts Options) (*types.AssessmentResult, error) {
	result := &types.AssessmentResult{
		ServiceName:   serviceName,
		ServiceURL:    serviceURL,
		AssessedAt:    time.Now().UTC(),
		AutomatedOnly: opts.AutomatedOnly,
	}

	allChecks := map[string]types.CheckResult{}

	for _, c := range e.checkers {
		checks, err := c.Check(serviceURL)
		if err != nil {
			result.Errors = append(result.Errors, fmt.Sprintf("checker %s: %s", c.Name(), err.Error()))
			continue
		}
		for _, chk := range checks {
			allChecks[chk.ID] = chk
		}
	}

	for _, sp := range standard.Points {
		point := types.StandardPoint{
			Number:      sp.Number,
			Title:       sp.Title,
			Description: sp.Description,
		}

		for _, checkID := range sp.CheckIDs {
			if chk, ok := allChecks[checkID]; ok {
				if opts.AutomatedOnly && !chk.Automated {
					continue
				}
				point.Checks = append(point.Checks, chk)
			}
		}

		result.Points = append(result.Points, point)
	}

	scoring.ScoreAssessment(result)

	return result, nil
}
