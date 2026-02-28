# Implementation Notes

## Overview

This implements v1 of the GDS Service Standard assessment tool as a Go CLI. The tool fetches a live service URL and runs automated checks against technically-automatable GDS standard points, producing a graded report (A–F).

## Assumptions

- Design artifacts from the design branch were not merged into main. Implementation decisions were derived directly from the planning decisions recorded in issue comments.
- The `golang.org/x/net/html` package is used for HTML parsing as it is part of the official Go extended standard library and does not require a browser.
- No browser automation (Chromium/CDP) is included in v1; that is deferred until a headless/headed browser runtime can be provisioned in the execution environment.
- The go.sum file contains the expected hashes — dependencies must be fetched with `go mod download` before building.

## Source Layout

```
src/
  go.mod                         – Go module definition
  cmd/assess/main.go             – CLI entry point (cobra)
  internal/
    types/types.go               – Core types: CheckResult, StandardPoint, AssessmentResult, Grade
    checker/
      checker.go                 – Checker interface
      web.go                     – HTTP/TLS checks (standard points 4, 9)
      accessibility.go           – WCAG/accessibility checks (standard point 5)
      performance.go             – Response time and caching checks (standard point 14)
      privacy.go                 – Privacy/cookie/tracker checks (standard point 9)
    engine/engine.go             – Orchestrates all checkers, maps results to standard points
    scoring/scoring.go           – Weighted graded scoring (A–F)
    standard/points.go           – Definitions for all 14 GDS standard points with check ID mappings
    output/output.go             – JSON and text output formatters
tests/
  go.mod                         – Test module (replaces src with local path)
  types_test.go
  scoring_test.go
  web_checker_test.go
  accessibility_test.go
  output_test.go
```

## Checks Implemented (v1)

| Check ID | Standard Point | Description |
|---|---|---|
| web-https | 9 | URL uses HTTPS with TLS 1.2+ |
| web-https-redirect | 9 | HTTP redirects to HTTPS |
| web-hsts | 9 | HSTS header present |
| web-security-headers | 9 | CSP, X-Content-Type-Options, X-Frame-Options, Referrer-Policy |
| web-cookie-flags | 9 | Session cookies have Secure and HttpOnly flags |
| web-html-lang | 4, 5 | HTML lang attribute set |
| web-page-title | 4 | Page has a descriptive title |
| web-viewport-meta | 4 | Mobile-ready viewport meta tag |
| a11y-img-alt | 5 | Images have alt text |
| a11y-form-labels | 5 | Form inputs have associated labels |
| a11y-heading-hierarchy | 5 | Heading levels follow logical order |
| a11y-skip-link | 5 | Skip to main content link present |
| a11y-link-text | 5 | Links have descriptive text |
| perf-response-time | 14 | Page loads in under 2 seconds |
| perf-cache-headers | 14 | Cache-Control header present |
| perf-compression | 14 | Response uses gzip/Brotli compression |
| privacy-link | 9 | Privacy/cookie policy link in page |
| privacy-cookie-banner | 9 | Cookie consent banner present |
| privacy-trackers | 9 | No known third-party trackers loaded |

Standard points 1, 2, 3, 6, 7, 8, 10, 11, 12, 13 have no automated checks in v1 and are reported as N/A.

## Scoring

- Each check produces a score in the range [0.0, 1.0].
- Standard point score = mean of its check scores × 100.
- Overall score = weighted mean of standard point scores (points 5 and 9 carry weight 1.5; all others 1.0).
- Grades: A ≥ 90, B ≥ 75, C ≥ 60, D ≥ 40, F < 40.

## CLI Usage

```sh
cd src
go build -o assess ./cmd/assess
./assess "Apply for a passport" --url https://www.gov.uk/apply-renew-passport
./assess "Apply for a passport" --url https://www.gov.uk/apply-renew-passport --format json -o result.json
./assess save "Apply for a passport" --url https://www.gov.uk/apply-renew-passport --dir results
```

## Running Tests

```sh
cd tests
go test ./... -v
```

## Deferred for Later Iterations

- Browser automation (Chromium/CDP) for JavaScript-rendered pages and interactive form flows.
- Authentication flows (with dummy credentials as noted in planning).
- Checking GOV.UK Design System usage (requires pattern matching against component library).
- Open-source repository check (standard point 12).
- Performance dashboard check (standard point 10).
- Parallel multi-checker execution.
- Database persistence (currently JSON file only).
