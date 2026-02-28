package checker

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/types"
	"golang.org/x/net/html"
)

// WebChecker performs HTTP-level checks against a service.
type WebChecker struct {
	client *http.Client
}

// CheckResultForTest is a type alias used in tests to access unexported fields.
type CheckResultForTest = types.CheckResult

// NewWebChecker returns a WebChecker with a 15-second timeout.
func NewWebChecker() *WebChecker {
	return &WebChecker{
		client: &http.Client{
			Timeout: 15 * time.Second,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				if len(via) >= 5 {
					return fmt.Errorf("too many redirects")
				}
				return nil
			},
		},
	}
}

// NewWebCheckerWithTransport returns a WebChecker using the provided transport (useful for tests).
func NewWebCheckerWithTransport(transport http.RoundTripper) *WebChecker {
	return &WebChecker{
		client: &http.Client{
			Timeout:   15 * time.Second,
			Transport: transport,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				if len(via) >= 5 {
					return fmt.Errorf("too many redirects")
				}
				return nil
			},
		},
	}
}

func (w *WebChecker) Name() string { return "web" }

func (w *WebChecker) Check(serviceURL string) ([]types.CheckResult, error) {
	var results []types.CheckResult

	results = append(results, w.checkHTTPS(serviceURL))
	results = append(results, w.checkSecurityHeaders(serviceURL))
	results = append(results, w.checkCookieFlags(serviceURL))
	results = append(results, w.checkHTTPSRedirect(serviceURL))
	results = append(results, w.checkHSTSHeader(serviceURL))

	htmlResults, err := w.checkHTMLContent(serviceURL)
	if err == nil {
		results = append(results, htmlResults...)
	}

	return results, nil
}

func (w *WebChecker) checkHTTPS(serviceURL string) types.CheckResult {
	r := types.CheckResult{
		ID:        "web-https",
		Name:      "HTTPS enforced",
		Automated: true,
	}

	if !strings.HasPrefix(serviceURL, "https://") {
		r.Passed = false
		r.Score = 0
		r.Evidence = "URL does not use HTTPS"
		r.Remediation = "Serve the service over HTTPS and redirect all HTTP traffic."
		return r
	}

	resp, err := w.client.Get(serviceURL)
	if err != nil {
		r.Error = err.Error()
		r.Score = 0
		return r
	}
	defer resp.Body.Close()

	if resp.TLS == nil {
		r.Passed = false
		r.Score = 0
		r.Evidence = "No TLS connection established"
		r.Remediation = "Ensure the server terminates TLS correctly."
		return r
	}

	switch resp.TLS.Version {
	case tls.VersionTLS13, tls.VersionTLS12:
		r.Passed = true
		r.Score = 1.0
		r.Evidence = fmt.Sprintf("TLS version: %s", tlsVersionName(resp.TLS.Version))
	default:
		r.Passed = false
		r.Score = 0.3
		r.Evidence = fmt.Sprintf("Outdated TLS version: %s", tlsVersionName(resp.TLS.Version))
		r.Remediation = "Upgrade to TLS 1.2 or TLS 1.3."
	}
	return r
}

func tlsVersionName(v uint16) string {
	switch v {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return fmt.Sprintf("unknown (0x%x)", v)
	}
}

func (w *WebChecker) checkSecurityHeaders(serviceURL string) types.CheckResult {
	r := types.CheckResult{
		ID:        "web-security-headers",
		Name:      "Security headers present",
		Automated: true,
	}

	resp, err := w.client.Get(serviceURL)
	if err != nil {
		r.Error = err.Error()
		return r
	}
	defer resp.Body.Close()

	wanted := []string{
		"Content-Security-Policy",
		"X-Content-Type-Options",
		"X-Frame-Options",
		"Referrer-Policy",
	}

	missing := []string{}
	found := []string{}
	for _, h := range wanted {
		if resp.Header.Get(h) != "" {
			found = append(found, h)
		} else {
			missing = append(missing, h)
		}
	}

	r.Score = float64(len(found)) / float64(len(wanted))
	r.Passed = len(missing) == 0
	r.Evidence = fmt.Sprintf("Present: %s. Missing: %s", strings.Join(found, ", "), strings.Join(missing, ", "))
	if len(missing) > 0 {
		r.Remediation = fmt.Sprintf("Add the following security headers: %s", strings.Join(missing, ", "))
	}
	return r
}

func (w *WebChecker) checkHSTSHeader(serviceURL string) types.CheckResult {
	r := types.CheckResult{
		ID:        "web-hsts",
		Name:      "HSTS header present",
		Automated: true,
	}

	resp, err := w.client.Get(serviceURL)
	if err != nil {
		r.Error = err.Error()
		return r
	}
	defer resp.Body.Close()

	hsts := resp.Header.Get("Strict-Transport-Security")
	if hsts == "" {
		r.Passed = false
		r.Score = 0
		r.Evidence = "Strict-Transport-Security header not present"
		r.Remediation = "Add a Strict-Transport-Security header with a long max-age (e.g. max-age=31536000; includeSubDomains; preload)."
		return r
	}

	r.Passed = true
	r.Score = 1.0
	r.Evidence = fmt.Sprintf("Strict-Transport-Security: %s", hsts)
	return r
}

func (w *WebChecker) checkCookieFlags(serviceURL string) types.CheckResult {
	r := types.CheckResult{
		ID:        "web-cookie-flags",
		Name:      "Cookies have Secure and HttpOnly flags",
		Automated: true,
	}

	resp, err := w.client.Get(serviceURL)
	if err != nil {
		r.Error = err.Error()
		return r
	}
	defer resp.Body.Close()

	cookies := resp.Cookies()
	if len(cookies) == 0 {
		r.Passed = true
		r.Score = 1.0
		r.Evidence = "No cookies set on this endpoint"
		return r
	}

	insecure := []string{}
	for _, c := range cookies {
		if !c.Secure || !c.HttpOnly {
			insecure = append(insecure, c.Name)
		}
	}

	if len(insecure) == 0 {
		r.Passed = true
		r.Score = 1.0
		r.Evidence = fmt.Sprintf("All %d cookie(s) have Secure and HttpOnly flags", len(cookies))
	} else {
		r.Passed = false
		r.Score = float64(len(cookies)-len(insecure)) / float64(len(cookies))
		r.Evidence = fmt.Sprintf("Cookies missing flags: %s", strings.Join(insecure, ", "))
		r.Remediation = "Set Secure and HttpOnly flags on all session cookies."
	}
	return r
}

func (w *WebChecker) checkHTTPSRedirect(serviceURL string) types.CheckResult {
	r := types.CheckResult{
		ID:        "web-https-redirect",
		Name:      "HTTP redirects to HTTPS",
		Automated: true,
	}

	if !strings.HasPrefix(serviceURL, "https://") {
		r.Passed = false
		r.Score = 0
		r.Evidence = "Service URL is not HTTPS"
		return r
	}

	httpURL := "http://" + strings.TrimPrefix(serviceURL, "https://")

	noRedirectClient := &http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := noRedirectClient.Get(httpURL)
	if err != nil {
		r.Error = err.Error()
		return r
	}
	defer resp.Body.Close()

	location := resp.Header.Get("Location")
	if (resp.StatusCode == 301 || resp.StatusCode == 302 || resp.StatusCode == 307 || resp.StatusCode == 308) &&
		strings.HasPrefix(location, "https://") {
		r.Passed = true
		r.Score = 1.0
		r.Evidence = fmt.Sprintf("HTTP returns %d to %s", resp.StatusCode, location)
	} else {
		r.Passed = false
		r.Score = 0
		r.Evidence = fmt.Sprintf("HTTP returns %d (Location: %s)", resp.StatusCode, location)
		r.Remediation = "Configure the server to redirect all HTTP requests to HTTPS with a 301 or 308 status."
	}
	return r
}

func (w *WebChecker) checkHTMLContent(serviceURL string) ([]types.CheckResult, error) {
	resp, err := w.client.Get(serviceURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	meta := extractMetaTags(doc)
	title := extractTitle(doc)
	langAttr := extractHTMLLang(doc)

	var results []types.CheckResult

	langResult := types.CheckResult{
		ID:        "web-html-lang",
		Name:      "HTML lang attribute set",
		Automated: true,
	}
	if langAttr != "" {
		langResult.Passed = true
		langResult.Score = 1.0
		langResult.Evidence = fmt.Sprintf("lang=%q", langAttr)
	} else {
		langResult.Passed = false
		langResult.Score = 0
		langResult.Evidence = "No lang attribute on <html> element"
		langResult.Remediation = "Add a lang attribute to the <html> element (e.g. lang=\"en\")."
	}
	results = append(results, langResult)

	titleResult := types.CheckResult{
		ID:        "web-page-title",
		Name:      "Page has a descriptive title",
		Automated: true,
	}
	if title != "" {
		titleResult.Passed = true
		titleResult.Score = 1.0
		titleResult.Evidence = fmt.Sprintf("title: %q", title)
	} else {
		titleResult.Passed = false
		titleResult.Score = 0
		titleResult.Evidence = "No <title> element found"
		titleResult.Remediation = "Add a descriptive <title> element to the page."
	}
	results = append(results, titleResult)

	viewportResult := types.CheckResult{
		ID:        "web-viewport-meta",
		Name:      "Viewport meta tag present (mobile-ready)",
		Automated: true,
	}
	viewport, hasViewport := meta["viewport"]
	if hasViewport && strings.Contains(viewport, "width=device-width") {
		viewportResult.Passed = true
		viewportResult.Score = 1.0
		viewportResult.Evidence = fmt.Sprintf("viewport: %q", viewport)
	} else if hasViewport {
		viewportResult.Passed = false
		viewportResult.Score = 0.5
		viewportResult.Evidence = fmt.Sprintf("viewport meta present but may not be mobile-optimised: %q", viewport)
		viewportResult.Remediation = "Set viewport to width=device-width, initial-scale=1."
	} else {
		viewportResult.Passed = false
		viewportResult.Score = 0
		viewportResult.Evidence = "No viewport meta tag found"
		viewportResult.Remediation = "Add <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">."
	}
	results = append(results, viewportResult)

	return results, nil
}

func extractMetaTags(n *html.Node) map[string]string {
	meta := map[string]string{}
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "meta" {
			var name, content string
			for _, a := range n.Attr {
				switch a.Key {
				case "name":
					name = strings.ToLower(a.Val)
				case "content":
					content = a.Val
				}
			}
			if name != "" {
				meta[name] = content
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(n)
	return meta
}

func extractTitle(n *html.Node) string {
	var title string
	var walk func(*html.Node) bool
	walk = func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "title" {
			if n.FirstChild != nil {
				title = n.FirstChild.Data
				return true
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if walk(c) {
				return true
			}
		}
		return false
	}
	walk(n)
	return title
}

func extractHTMLLang(n *html.Node) string {
	var walk func(*html.Node) string
	walk = func(n *html.Node) string {
		if n.Type == html.ElementNode && n.Data == "html" {
			for _, a := range n.Attr {
				if a.Key == "lang" {
					return a.Val
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if v := walk(c); v != "" {
				return v
			}
		}
		return ""
	}
	return walk(n)
}
