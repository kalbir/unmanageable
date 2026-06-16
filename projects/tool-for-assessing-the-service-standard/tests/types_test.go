package tests

import (
	"testing"

	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/types"
)

func TestGradeFromScore(t *testing.T) {
	cases := []struct {
		score float64
		want  types.Grade
	}{
		{100, types.GradeA},
		{90, types.GradeA},
		{89.9, types.GradeB},
		{75, types.GradeB},
		{74.9, types.GradeC},
		{60, types.GradeC},
		{59.9, types.GradeD},
		{40, types.GradeD},
		{39.9, types.GradeF},
		{0, types.GradeF},
	}

	for _, tc := range cases {
		got := types.GradeFromScore(tc.score)
		if got != tc.want {
			t.Errorf("GradeFromScore(%.1f) = %s, want %s", tc.score, got, tc.want)
		}
	}
}
