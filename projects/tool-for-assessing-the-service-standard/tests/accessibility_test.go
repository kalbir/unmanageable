package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/checker"
)

func TestAccessibilityChecker_AllGood(t *testing.T) {
	page := `<!DOCTYPE html>
<html lang="en">
<head><title>Good Service</title></head>
<body>
  <a href="#main">Skip to main content</a>
  <main id="main">
    <h1>Apply for something</h1>
    <h2>Section one</h2>
    <img src="logo.png" alt="Logo">
    <form>
      <label for="name">Name</label>
      <input id="name" type="text">
    </form>
    <a href="/about">About this service</a>
  </main>
</body>
</html>`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(page))
	}))
	defer srv.Close()

	ac := checker.NewAccessibilityChecker()
	results, err := ac.Check(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	checkMap := map[string]bool{}
	for _, r := range results {
		checkMap[r.ID] = r.Passed
	}

	if !checkMap["a11y-img-alt"] {
		t.Error("expected a11y-img-alt to pass")
	}
	if !checkMap["a11y-form-labels"] {
		t.Error("expected a11y-form-labels to pass")
	}
	if !checkMap["a11y-skip-link"] {
		t.Error("expected a11y-skip-link to pass")
	}
	if !checkMap["a11y-heading-hierarchy"] {
		t.Error("expected a11y-heading-hierarchy to pass")
	}
}

func TestAccessibilityChecker_MissingAltText(t *testing.T) {
	page := `<!DOCTYPE html>
<html lang="en"><head><title>Bad</title></head>
<body><h1>Test</h1><img src="photo.jpg"></body>
</html>`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(page))
	}))
	defer srv.Close()

	ac := checker.NewAccessibilityChecker()
	results, err := ac.Check(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	imgAlt := findCheck(results, "a11y-img-alt")
	if imgAlt == nil {
		t.Fatal("a11y-img-alt check not found")
	}
	if imgAlt.Passed {
		t.Error("expected a11y-img-alt to fail when img has no alt")
	}
	if imgAlt.Remediation == "" {
		t.Error("expected remediation guidance for missing alt text")
	}
}

func TestAccessibilityChecker_SkippedHeadings(t *testing.T) {
	page := `<!DOCTYPE html>
<html lang="en"><head><title>Bad Headings</title></head>
<body>
  <h1>Title</h1>
  <h3>Skipped h2</h3>
</body>
</html>`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(page))
	}))
	defer srv.Close()

	ac := checker.NewAccessibilityChecker()
	results, err := ac.Check(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	headings := findCheck(results, "a11y-heading-hierarchy")
	if headings == nil {
		t.Fatal("a11y-heading-hierarchy check not found")
	}
	if headings.Passed {
		t.Error("expected a11y-heading-hierarchy to fail for skipped levels")
	}
}
