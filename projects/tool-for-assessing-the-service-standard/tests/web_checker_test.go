package tests

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/checker"
	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/types"
)

func findCheck(results []types.CheckResult, id string) *types.CheckResult {
	for i := range results {
		if results[i].ID == id {
			return &results[i]
		}
	}
	return nil
}

func TestWebChecker_HTTPSCheck_HTTP(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	defer srv.Close()

	wc := checker.NewWebChecker()
	results, err := wc.Check(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	httpsCheck := findCheck(results, "web-https")
	if httpsCheck == nil {
		t.Fatal("web-https check not found in results")
	}
	if httpsCheck.Passed {
		t.Error("expected web-https to fail for HTTP server")
	}
}

func TestWebChecker_HTTPSCheck_TLS(t *testing.T) {
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000")
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`<!DOCTYPE html><html lang="en"><head><title>Test</title><meta name="viewport" content="width=device-width, initial-scale=1"></head><body></body></html>`))
	}))
	defer srv.Close()

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	wc := checker.NewWebCheckerWithTransport(transport)
	results, err := wc.Check(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	checkMap := map[string]bool{}
	for _, r := range results {
		checkMap[r.ID] = r.Passed
	}

	if !checkMap["web-hsts"] {
		t.Error("expected web-hsts to pass when HSTS header is present")
	}
	if !checkMap["web-security-headers"] {
		t.Error("expected web-security-headers to pass when all headers present")
	}
}

func TestWebChecker_MissingSecurityHeaders(t *testing.T) {
	srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte(`<html><body></body></html>`))
	}))
	defer srv.Close()

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	wc := checker.NewWebCheckerWithTransport(transport)
	results, err := wc.Check(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	secHeaders := findCheck(results, "web-security-headers")
	if secHeaders != nil && secHeaders.Passed {
		t.Error("expected web-security-headers to fail when headers are missing")
	}

	hsts := findCheck(results, "web-hsts")
	if hsts != nil && hsts.Passed {
		t.Error("expected web-hsts to fail when HSTS header is missing")
	}
}
