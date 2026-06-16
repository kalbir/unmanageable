package checker

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/types"
	"golang.org/x/net/html"
)

// PrivacyChecker checks for privacy-related GDS requirements (standard point 9).
type PrivacyChecker struct {
	client *http.Client
}

// NewPrivacyChecker returns a PrivacyChecker.
func NewPrivacyChecker() *PrivacyChecker {
	return &PrivacyChecker{
		client: &http.Client{Timeout: 15 * time.Second},
	}
}

func (p *PrivacyChecker) Name() string { return "privacy" }

func (p *PrivacyChecker) Check(serviceURL string) ([]types.CheckResult, error) {
	var results []types.CheckResult

	results = append(results, p.checkPrivacyLink(serviceURL))
	results = append(results, p.checkCookieBanner(serviceURL))
	results = append(results, p.checkThirdPartyTrackers(serviceURL))

	return results, nil
}

func (p *PrivacyChecker) checkPrivacyLink(serviceURL string) types.CheckResult {
	r := types.CheckResult{
		ID:        "privacy-link",
		Name:      "Privacy notice link present",
		Automated: true,
	}

	resp, err := p.client.Get(serviceURL)
	if err != nil {
		r.Error = err.Error()
		return r
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		r.Error = err.Error()
		return r
	}

	found := false
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			text := strings.ToLower(nodeText(n))
			href := ""
			for _, a := range n.Attr {
				if a.Key == "href" {
					href = strings.ToLower(a.Val)
				}
			}
			if strings.Contains(text, "privacy") || strings.Contains(href, "privacy") ||
				strings.Contains(text, "cookie") || strings.Contains(href, "cookie") {
				found = true
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)

	if found {
		r.Passed = true
		r.Score = 1.0
		r.Evidence = "Privacy or cookie policy link detected"
	} else {
		r.Passed = false
		r.Score = 0
		r.Evidence = "No privacy or cookie policy link found"
		r.Remediation = "Add a link to your privacy notice and cookie policy, typically in the footer."
	}
	return r
}

func (p *PrivacyChecker) checkCookieBanner(serviceURL string) types.CheckResult {
	r := types.CheckResult{
		ID:        "privacy-cookie-banner",
		Name:      "Cookie consent banner present",
		Automated: true,
	}

	resp, err := p.client.Get(serviceURL)
	if err != nil {
		r.Error = err.Error()
		return r
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		r.Error = err.Error()
		return r
	}

	hasCookies := len(resp.Cookies()) > 0
	bannerKeywords := []string{"cookie", "consent", "accept cookies", "analytics"}
	found := false

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" || a.Key == "class" {
					lower := strings.ToLower(a.Val)
					for _, kw := range bannerKeywords {
						if strings.Contains(lower, kw) {
							found = true
						}
					}
				}
			}
			text := strings.ToLower(nodeText(n))
			for _, kw := range bannerKeywords {
				if strings.Contains(text, kw+" ") {
					found = true
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)

	if found {
		r.Passed = true
		r.Score = 1.0
		r.Evidence = "Cookie consent element detected in page"
	} else if !hasCookies {
		r.Passed = true
		r.Score = 0.8
		r.Evidence = "No cookies set and no banner found (may not be needed)"
	} else {
		r.Passed = false
		r.Score = 0
		r.Evidence = "Cookies are set but no cookie consent banner detected"
		r.Remediation = "Implement a cookie consent banner and only set non-essential cookies after user consent."
	}
	return r
}

func (p *PrivacyChecker) checkThirdPartyTrackers(serviceURL string) types.CheckResult {
	r := types.CheckResult{
		ID:        "privacy-trackers",
		Name:      "No known third-party trackers loaded",
		Automated: true,
	}

	resp, err := p.client.Get(serviceURL)
	if err != nil {
		r.Error = err.Error()
		return r
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		r.Error = err.Error()
		return r
	}

	knownTrackers := []string{
		"google-analytics.com",
		"googletagmanager.com",
		"facebook.com/tr",
		"doubleclick.net",
		"hotjar.com",
		"fullstory.com",
		"mixpanel.com",
	}

	found := []string{}
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "img" || n.Data == "iframe") {
			for _, a := range n.Attr {
				if a.Key == "src" || a.Key == "href" {
					for _, tracker := range knownTrackers {
						if strings.Contains(strings.ToLower(a.Val), tracker) {
							found = append(found, tracker)
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)

	if len(found) == 0 {
		r.Passed = true
		r.Score = 1.0
		r.Evidence = "No known third-party trackers detected in page source"
	} else {
		r.Passed = false
		r.Score = 0
		r.Evidence = fmt.Sprintf("Third-party trackers detected: %s", strings.Join(found, ", "))
		r.Remediation = "Only load third-party analytics or trackers after explicit user consent. Consider using privacy-focused alternatives."
	}
	return r
}
