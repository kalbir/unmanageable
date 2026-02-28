package checker

import (
	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/types"
)

// Checker is implemented by each assessment plugin.
type Checker interface {
	Name() string
	Check(serviceURL string) ([]types.CheckResult, error)
}
