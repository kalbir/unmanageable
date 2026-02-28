package checker

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/kalbir/unmanageable/projects/tool-for-assessing-the-service-standard/internal/types"
	"golang.org/x/net/html"
)

// AccessibilityChecker checks for basic WCAG/GDS accessibility requirements (standard point 5).
type AccessibilityChecker struct {
	client *http.Client
}

// NewAccessibilityChecker returns a new AccessibilityChecker.
func NewAccessibilityChecker() *AccessibilityChecker {
	return &AccessibilityChecker{
		client: &http.Client{Timeout: 15 * time.Second},
	}
}

func (a *AccessibilityChecker) Name() string { return "accessibility" }

func (a *AccessibilityChecker) Check(serviceURL string) ([]types.CheckResult, error) {
	resp, err := a.client.Get(serviceURL)
	if err != nil {
		return []types.CheckResult{{
			ID:    "accessibility-fetch",
			Name:  "Fetch page for accessibility checks",
			Error: err.Error(),
		}}, nil
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing HTML: %w", err)
	}

	var results []types.CheckResult
	results = append(results, checkImgAltText(doc))
	results = append(results, checkFormLabels(doc))
	results = append(results, checkHeadingHierarchy(doc))
	results = append(results, checkSkipLink(doc))
	results = append(results, checkLinkText(doc))
	return results, nil
}

func checkImgAltText(doc *html.Node) types.CheckResult {
	r := types.CheckResult{
		ID:        "a11y-img-alt",
		Name:      "Images have alt text",
		Automated: true,
	}

	total, missing := 0, 0
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			total++
			hasAlt := false
			for _, a := range n.Attr {
				if a.Key == "alt" {
					hasAlt = true
					break
				}
			}
			if !hasAlt {
				missing++
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)

	if total == 0 {
		r.Passed = true
		r.Score = 1.0
		r.Evidence = "No <img> elements found"
		return r
	}

	r.Score = float64(total-missing) / float64(total)
	r.Passed = missing == 0
	r.Evidence = fmt.Sprintf("%d of %d images have alt text", total-missing, total)
	if missing > 0 {
		r.Remediation = "Add descriptive alt attributes to all <img> elements. Use alt=\"\" for decorative images."
	}
	return r
}

func checkFormLabels(doc *html.Node) types.CheckResult {
	r := types.CheckResult{
		ID:        "a11y-form-labels",
		Name:      "Form inputs have associated labels",
		Automated: true,
	}

	inputs, labeled := 0, 0
	inputIDs := map[string]bool{}

	var collectLabelFor func(*html.Node)
	collectLabelFor = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "label" {
			for _, a := range n.Attr {
				if a.Key == "for" {
					inputIDs[a.Val] = true
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			collectLabelFor(c)
		}
	}
	collectLabelFor(doc)

	var walkInputs func(*html.Node)
	walkInputs = func(n *html.Node) {
		if n.Type == html.ElementNode && (n.Data == "input" || n.Data == "select" || n.Data == "textarea") {
			inputType := ""
			inputID := ""
			ariaLabel := ""
			for _, a := range n.Attr {
				switch a.Key {
				case "type":
					inputType = a.Val
				case "id":
					inputID = a.Val
				case "aria-label", "aria-labelledby":
					ariaLabel = a.Val
				}
			}
			if inputType == "hidden" || inputType == "submit" || inputType == "button" || inputType == "reset" {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					walkInputs(c)
				}
				return
			}
			inputs++
			if (inputID != "" && inputIDs[inputID]) || ariaLabel != "" {
				labeled++
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walkInputs(c)
		}
	}
	walkInputs(doc)

	if inputs == 0 {
		r.Passed = true
		r.Score = 1.0
		r.Evidence = "No form inputs found"
		return r
	}

	r.Score = float64(labeled) / float64(inputs)
	r.Passed = labeled == inputs
	r.Evidence = fmt.Sprintf("%d of %d inputs have labels", labeled, inputs)
	if labeled < inputs {
		r.Remediation = "Associate every form input with a <label> element using the 'for' attribute, or add an aria-label."
	}
	return r
}

func checkHeadingHierarchy(doc *html.Node) types.CheckResult {
	r := types.CheckResult{
		ID:        "a11y-heading-hierarchy",
		Name:      "Headings follow a logical hierarchy",
		Automated: true,
	}

	var headings []int
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "h1", "h2", "h3", "h4", "h5", "h6":
				level := int(n.Data[1] - '0')
				headings = append(headings, level)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)

	if len(headings) == 0 {
		r.Passed = false
		r.Score = 0
		r.Evidence = "No heading elements found"
		r.Remediation = "Use heading elements (h1–h6) to structure page content."
		return r
	}

	skips := 0
	for i := 1; i < len(headings); i++ {
		if headings[i]-headings[i-1] > 1 {
			skips++
		}
	}

	hasH1 := false
	for _, h := range headings {
		if h == 1 {
			hasH1 = true
			break
		}
	}

	issues := skips
	if !hasH1 {
		issues++
	}

	if issues == 0 {
		r.Passed = true
		r.Score = 1.0
		r.Evidence = fmt.Sprintf("%d headings, hierarchy is correct", len(headings))
	} else {
		r.Passed = false
		r.Score = 0.5
		evidenceParts := []string{}
		if !hasH1 {
			evidenceParts = append(evidenceParts, "no h1 found")
		}
		if skips > 0 {
			evidenceParts = append(evidenceParts, fmt.Sprintf("%d heading level skip(s)", skips))
		}
		r.Evidence = strings.Join(evidenceParts, "; ")
		r.Remediation = "Ensure the page has a single h1 and that heading levels are not skipped (e.g. h2 → h4 skips h3)."
	}
	return r
}

func checkSkipLink(doc *html.Node) types.CheckResult {
	r := types.CheckResult{
		ID:        "a11y-skip-link",
		Name:      "Skip to main content link present",
		Automated: true,
	}

	found := false
	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			text := nodeText(n)
			href := ""
			for _, a := range n.Attr {
				if a.Key == "href" {
					href = a.Val
				}
			}
			lower := strings.ToLower(text)
			if (strings.Contains(lower, "skip") || strings.Contains(lower, "jump")) &&
				strings.HasPrefix(href, "#") {
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
		r.Evidence = "Skip to main content link detected"
	} else {
		r.Passed = false
		r.Score = 0
		r.Evidence = "No skip navigation link found"
		r.Remediation = "Add a visible or visually-hidden 'Skip to main content' link as the first focusable element."
	}
	return r
}

func checkLinkText(doc *html.Node) types.CheckResult {
	r := types.CheckResult{
		ID:        "a11y-link-text",
		Name:      "Links have descriptive text",
		Automated: true,
	}

	nonDescriptive := []string{"click here", "here", "read more", "more", "link", "this"}
	total, bad := 0, 0

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			text := strings.ToLower(strings.TrimSpace(nodeText(n)))
			if text == "" {
				for _, a := range n.Attr {
					if a.Key == "aria-label" && a.Val != "" {
						text = strings.ToLower(a.Val)
					}
				}
			}
			if text != "" {
				total++
				for _, nd := range nonDescriptive {
					if text == nd {
						bad++
						break
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}
	walk(doc)

	if total == 0 {
		r.Passed = true
		r.Score = 1.0
		r.Evidence = "No anchor elements found"
		return r
	}

	r.Score = float64(total-bad) / float64(total)
	r.Passed = bad == 0
	r.Evidence = fmt.Sprintf("%d of %d links have descriptive text", total-bad, total)
	if bad > 0 {
		r.Remediation = "Replace non-descriptive link text such as 'click here' or 'read more' with text that describes the destination."
	}
	return r
}

func nodeText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var sb strings.Builder
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		sb.WriteString(nodeText(c))
	}
	return sb.String()
}
