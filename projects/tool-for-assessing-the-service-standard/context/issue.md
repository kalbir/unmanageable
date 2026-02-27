# Issue #2: Tool for assessing the service standard

## Original Issue

I want to build an application that, given the name of a government service will assess in real time
whether that service meets the published GDS service standard. Not by using published data sets but by using real time, potentially agentic, capabilities to test the service.

## Context

The GDS (Government Digital Service) Service Standard is a set of 14 criteria that all public-facing UK government services must meet. The standard covers areas including:

1. Understanding users and their needs
2. Solving a whole problem for users
3. Providing a joined up experience across all channels
4. Making the service simple to use
5. Making sure everyone can use the service
6. Having a multidisciplinary team
7. Using agile ways of working
8. Iterating and improving frequently
9. Creating a secure service which protects users' privacy
10. Defining what success looks like and publishing performance data
11. Choosing the right tools and technology
12. Making new source code open
13. Using and contributing to open standards, common components and patterns
14. Operating a reliable service

## Requirements

- **Input**: Government service name or URL
- **Processing**: Real-time automated assessment against the service standard
- **Output**: Comprehensive report on compliance with each standard point
- **Approach**: Agentic/automated testing rather than static data analysis

## Technical Approach

The tool will employ multiple assessment techniques:
- Web scraping and content analysis
- Accessibility testing (WCAG compliance)
- Security scanning (headers, HTTPS, vulnerabilities)
- Performance testing (Core Web Vitals, load times)
- Pattern recognition for design compliance
- API discovery and testing
- Natural language processing for content evaluation

## Success Criteria

- Ability to assess any publicly accessible UK government service
- Automated evaluation of technically measurable standard points
- Clear reporting of compliance status with evidence
- Actionable recommendations for non-compliant areas
- Real-time assessment without relying on pre-published datasets


