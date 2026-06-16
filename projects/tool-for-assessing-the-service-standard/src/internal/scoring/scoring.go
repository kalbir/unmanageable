package scoring

import (
	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/types"
)

// pointWeights assigns relative weight to each GDS standard point number.
var pointWeights = map[int]float64{
	1:  1.0,
	2:  1.0,
	3:  1.0,
	4:  1.0,
	5:  1.5,
	6:  1.0,
	7:  1.0,
	8:  1.0,
	9:  1.5,
	10: 1.0,
	11: 1.0,
	12: 1.0,
	13: 1.0,
	14: 1.0,
}

// ScorePoint computes the grade and aggregate score for a single standard point
// based on the check results it contains.
func ScorePoint(point *types.StandardPoint) {
	if len(point.Checks) == 0 {
		point.Score = 0
		point.Grade = types.GradeF
		return
	}

	var total, sum float64
	for _, c := range point.Checks {
		if c.Error != "" {
			continue
		}
		total++
		sum += c.Score
	}

	if total == 0 {
		point.Score = 0
		point.Grade = types.GradeF
		return
	}

	point.Score = (sum / total) * 100
	point.Grade = types.GradeFromScore(point.Score)
}

// ScoreAssessment computes the overall grade and score across all standard points.
func ScoreAssessment(result *types.AssessmentResult) {
	if len(result.Points) == 0 {
		result.OverallScore = 0
		result.OverallGrade = types.GradeF
		return
	}

	var weightedSum, totalWeight float64
	for i := range result.Points {
		pt := &result.Points[i]
		ScorePoint(pt)
		w := pointWeights[pt.Number]
		if w == 0 {
			w = 1.0
		}
		weightedSum += pt.Score * w
		totalWeight += w
	}

	result.OverallScore = weightedSum / totalWeight
	result.OverallGrade = types.GradeFromScore(result.OverallScore)
}
