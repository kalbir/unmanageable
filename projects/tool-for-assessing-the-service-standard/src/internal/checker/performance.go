package checker

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/types"
)

// PerformanceChecker checks response time and basic caching headers (standard point 14).
type PerformanceChecker struct {
	client *http.Client
}

// NewPerformanceChecker returns a PerformanceChecker.
func NewPerformanceChecker() *PerformanceChecker {
	return &PerformanceChecker{
		client: &http.Client{Timeout: 30 * time.Second},
	}
}

func (p *PerformanceChecker) Name() string { return "performance" }

func (p *PerformanceChecker) Check(serviceURL string) ([]types.CheckResult, error) {
	var results []types.CheckResult

	results = append(results, p.checkResponseTime(serviceURL))
	results = append(results, p.checkCacheHeaders(serviceURL))
	results = append(results, p.checkCompression(serviceURL))

	return results, nil
}

func (p *PerformanceChecker) checkResponseTime(serviceURL string) types.CheckResult {
	r := types.CheckResult{
		ID:        "perf-response-time",
		Name:      "Page response time under 2 seconds",
		Automated: true,
	}

	start := time.Now()
	resp, err := p.client.Get(serviceURL)
	elapsed := time.Since(start)

	if err != nil {
		r.Error = err.Error()
		return r
	}
	defer resp.Body.Close()

	ms := elapsed.Milliseconds()
	r.Evidence = fmt.Sprintf("Response time: %dms", ms)

	switch {
	case ms < 500:
		r.Passed = true
		r.Score = 1.0
	case ms < 1000:
		r.Passed = true
		r.Score = 0.85
	case ms < 2000:
		r.Passed = true
		r.Score = 0.7
	case ms < 4000:
		r.Passed = false
		r.Score = 0.4
		r.Remediation = "Response time exceeds 2 seconds. Investigate server-side performance, caching, and CDN usage."
	default:
		r.Passed = false
		r.Score = 0.1
		r.Remediation = "Response time exceeds 4 seconds. Urgent performance investigation required."
	}
	return r
}

func (p *PerformanceChecker) checkCacheHeaders(serviceURL string) types.CheckResult {
	r := types.CheckResult{
		ID:        "perf-cache-headers",
		Name:      "Cache-Control header present",
		Automated: true,
	}

	resp, err := p.client.Get(serviceURL)
	if err != nil {
		r.Error = err.Error()
		return r
	}
	defer resp.Body.Close()

	cacheControl := resp.Header.Get("Cache-Control")
	if cacheControl != "" {
		r.Passed = true
		r.Score = 1.0
		r.Evidence = fmt.Sprintf("Cache-Control: %s", cacheControl)
	} else {
		r.Passed = false
		r.Score = 0
		r.Evidence = "No Cache-Control header"
		r.Remediation = "Set appropriate Cache-Control headers. Use 'no-store' for sensitive/dynamic pages and appropriate max-age for static assets."
	}
	return r
}

func (p *PerformanceChecker) checkCompression(serviceURL string) types.CheckResult {
	r := types.CheckResult{
		ID:        "perf-compression",
		Name:      "Response uses compression (gzip/br)",
		Automated: true,
	}

	req, err := http.NewRequest("GET", serviceURL, nil)
	if err != nil {
		r.Error = err.Error()
		return r
	}
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")

	resp, err := p.client.Do(req)
	if err != nil {
		r.Error = err.Error()
		return r
	}
	defer resp.Body.Close()

	encoding := resp.Header.Get("Content-Encoding")
	contentLength := resp.Header.Get("Content-Length")

	if encoding == "gzip" || encoding == "br" || encoding == "deflate" {
		r.Passed = true
		r.Score = 1.0
		r.Evidence = fmt.Sprintf("Content-Encoding: %s", encoding)
	} else if contentLength != "" {
		size, _ := strconv.ParseInt(contentLength, 10, 64)
		if size < 1024 {
			r.Passed = true
			r.Score = 0.8
			r.Evidence = fmt.Sprintf("No compression but response is small (%d bytes)", size)
		} else {
			r.Passed = false
			r.Score = 0.3
			r.Evidence = fmt.Sprintf("No compression, response size: %d bytes", size)
			r.Remediation = "Enable gzip or Brotli compression on the server to reduce transfer sizes."
		}
	} else {
		r.Passed = false
		r.Score = 0.3
		r.Evidence = "No Content-Encoding header and Content-Length unknown"
		r.Remediation = "Enable gzip or Brotli compression on the server."
	}
	return r
}
