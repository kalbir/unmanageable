package standard

// Point describes one of the 14 GDS service standard points.
type Point struct {
	Number      int
	Title       string
	Description string
	CheckIDs    []string
}

// Points is the list of all 14 GDS service standard points with the check IDs that are
// mapped to them. Only technically-automatable checks are included in v1.
var Points = []Point{
	{
		Number:      1,
		Title:       "Understand users and their needs",
		Description: "Develop a deep understanding of users and the problem you're trying to solve for them.",
		CheckIDs:    nil,
	},
	{
		Number:      2,
		Title:       "Solve a whole problem for users",
		Description: "Work towards creating a service that solves one whole problem for users, working with other teams and organisations where needed.",
		CheckIDs:    nil,
	},
	{
		Number:      3,
		Title:       "Provide a joined up experience across all channels",
		Description: "Work towards creating a service that meets users' needs across all channels, including online, phone, paper and face to face.",
		CheckIDs:    nil,
	},
	{
		Number:      4,
		Title:       "Make the service simple to use",
		Description: "Build a service that's simple, intuitive and works in a way that is familiar to users.",
		CheckIDs:    []string{"web-page-title", "web-viewport-meta", "web-html-lang", "a11y-heading-hierarchy"},
	},
	{
		Number:      5,
		Title:       "Make sure everyone can use the service",
		Description: "Provide a service that everyone can use, including disabled people and those who need assistance.",
		CheckIDs:    []string{"a11y-img-alt", "a11y-form-labels", "a11y-skip-link", "a11y-link-text", "web-html-lang"},
	},
	{
		Number:      6,
		Title:       "Have a multidisciplinary team",
		Description: "Put in place a multidisciplinary team that can create and operate the service.",
		CheckIDs:    nil,
	},
	{
		Number:      7,
		Title:       "Use agile ways of working",
		Description: "Create the service using agile, iterative, user-centred methods.",
		CheckIDs:    nil,
	},
	{
		Number:      8,
		Title:       "Iterate and improve frequently",
		Description: "Make sure you have the capacity, resources and technical flexibility to iterate and improve the service.",
		CheckIDs:    nil,
	},
	{
		Number:      9,
		Title:       "Create a secure service which protects users' privacy",
		Description: "Evaluate what data the service will be collecting, storing and providing, and address the security level and risks associated with the service.",
		CheckIDs:    []string{"web-https", "web-https-redirect", "web-hsts", "web-security-headers", "web-cookie-flags", "privacy-link", "privacy-cookie-banner", "privacy-trackers"},
	},
	{
		Number:      10,
		Title:       "Define what success looks like and publish performance data",
		Description: "Identify performance indicators for the service, including the 4 mandatory key performance indicators (KPIs) defined in the manual.",
		CheckIDs:    nil,
	},
	{
		Number:      11,
		Title:       "Choose the right tools and technology",
		Description: "Choose tools and technology that let you create a high quality service in a cost effective way.",
		CheckIDs:    nil,
	},
	{
		Number:      12,
		Title:       "Make new source code open",
		Description: "Make all new source code open and reusable, and publish it under appropriate licences.",
		CheckIDs:    nil,
	},
	{
		Number:      13,
		Title:       "Use and contribute to open standards, common components and patterns",
		Description: "Build on open standards and common components and patterns from the GOV.UK Design System where relevant.",
		CheckIDs:    nil,
	},
	{
		Number:      14,
		Title:       "Operate a reliable service",
		Description: "Minimise service downtime and have a plan to deal with it when it does happen.",
		CheckIDs:    []string{"perf-response-time", "perf-cache-headers", "perf-compression"},
	},
}
