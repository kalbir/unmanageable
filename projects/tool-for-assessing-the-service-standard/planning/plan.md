# Planning Document: GDS Service Standard Assessment Tool

## Overview

This tool will assess government services against the GDS (Government Digital Service) service standard in real-time. The assessment will use automated testing capabilities including web scraping, accessibility testing, security scanning, and potentially agentic approaches to evaluate services without relying on published datasets.

**Technology Stack**: Go
**Deployment Model**: CLI tool
**Scoring Method**: Graded scoring system
**Target Version**: v1 - Focus on technically feasible automations

## Tasks (Ordered)

1. **Research and Document GDS Standards**
   - Analyze all 14 GDS service standard points
   - Categorize which can be technically automated vs. requiring manual review
   - Define grading criteria for each assessable point

2. **Design Core Architecture**
   - CLI interface structure
   - Service discovery and entry point detection
   - Assessment module framework (pluggable for different tests)
   - Result storage system (file or database output)

3. **Implement Service Discovery Module**
   - Input: government service name
   - URL discovery and validation
   - Service type classification (web service focus for v1)
   - Entry point identification

4. **Build Assessment Engines**
   - Accessibility testing module (WCAG compliance)
   - Performance testing module
   - Security headers checker
   - Mobile responsiveness tester
   - Plain English content analyzer
   - User journey tracer (as far as possible without credentials)

5. **Develop Scoring System**
   - Graded scoring algorithm
   - Weight assignment per standard point
   - Partial automation handling (flagging for manual review)
   - Remediation guidance generation

6. **Create Output System**
   - Structured result format (JSON/structured text)
   - File-based storage for v1
   - Report generation with scores and recommendations
   - Optional manual review prompts

7. **Testing and Validation**
   - Unit tests for each assessment module
   - Integration tests with known government services
   - Validation of scoring accuracy
   - Edge case handling (service unavailable, partial failures)

## Dependencies

### Technical Dependencies
- **Go standard library** for core functionality
- **Colly or similar** for web scraping
- **Chrome/Chromium** for headless browser testing
- **Playwright or Selenium** for browser automation
- **axe-core** or similar for accessibility testing
- **Security scanning libraries** for header/SSL checks

### External Dependencies
- Public access to government service websites
- No requirement for service credentials (v1 limitation)
- Stable internet connection for real-time assessment

## Assumptions

1. Government services have public-facing web interfaces
2. Services can be identified by name and have discoverable URLs
3. Assessment can provide value without authenticated access
4. GDS service standard points remain stable during development
5. Automated testing is legally permissible on public government sites
6. Services follow standard web conventions (HTML, HTTP/HTTPS)

## Risks

### Technical Risks
- **Rate limiting**: Government sites may block automated testing
- **Dynamic content**: JavaScript-heavy sites may require complex browser automation
- **Service diversity**: Wide variation in service implementations
- **False positives/negatives**: Automated tests may not capture full context

### Operational Risks
- **Legal considerations**: Need to ensure testing is within acceptable use
- **Service changes**: Frequent updates to services may break assessments
- **Incomplete coverage**: Some standards cannot be fully automated

### Mitigation Strategies
- Implement polite crawling with delays
- Use headless browsers for JavaScript rendering
- Build modular system to handle diverse implementations
- Clear documentation of assessment limitations
- Include confidence scores with results

## Out of Scope

1. Authentication-required service areas (for v1)
2. Non-web services (mobile apps, phone services)
3. Internal government systems
4. Historical/trend analysis (v1 is single-point-in-time)
5. Automated fixes or patches
6. Integration with government systems
7. Multi-service batch processing (one at a time for v1)
8. Caching of results (v1 - may add later)
9. Web UI (CLI only for v1)
10. API endpoints (CLI only for v1)

## Resolved Decisions

Based on user input, the following decisions have been made:

### Technical Decisions
- **Language**: Go
- **Deployment**: CLI tool
- **Storage**: Results written to file or database (no caching in v1)
- **Authentication**: Get as far as possible without credentials, future iterations may use dummy credentials
- **Browser modes**: Both headless and headed browser capabilities

### Scope Decisions
- **v1 Focus**: Technically feasible automations only
- **Service depth**: As far as possible without credentials
- **Manual components**: No manual assessments in v1
- **Service types**: Web services only for v1
- **Output**: Provide remediation guidance

### Implementation Decisions
- **Scoring**: Graded scoring system
- **Partial automation**: Flag for manual review with optional mode/flag for fully automated only
- **Service processing**: One service at a time
- **Entry points**: Pick one entry point per service

### Policy Decisions
- **Legal restrictions**: None identified by user
- **Result visibility**: Public to tool runner, not published
- **Accessibility issues**: Note in tool output
- **Service communication**: None required (alpha version)
- **Government coordination**: Not required

## GDS Service Standard Points

The 14 points of the GDS Service Standard that will be assessed:

1. **Understand users and their needs** - Partially automatable (content analysis)
2. **Solve a whole problem for users** - Manual review required
3. **Provide a joined up experience** - Partially automatable (link analysis)
4. **Make the service simple to use** - Partially automatable (complexity metrics)
5. **Make sure everyone can use the service** - Automatable (accessibility testing)
6. **Have a multidisciplinary team** - Not automatable
7. **Use agile ways of working** - Not automatable
8. **Iterate and improve frequently** - Partially automatable (update frequency)
9. **Create a secure service** - Automatable (security headers, SSL)
10. **Define what success looks like** - Manual review required
11. **Choose the right tools and technology** - Partially automatable (tech stack analysis)
12. **Make new source code open** - Automatable (check for code repositories)
13. **Use and contribute to open standards** - Partially automatable
14. **Operate a reliable service** - Automatable (uptime, performance)

## Next Phase

With planning complete and decisions resolved, the next phase is **Design**, which will produce:
- Detailed system architecture
- Module interfaces and data flow
- CLI command structure
- Assessment plugin framework
- Scoring algorithm specification
- Output format schemas