# Planning Document: Tool for Assessing the Service Standard

## Overview

This project will build a tool that assesses government digital services against the GDS Service Standard in real-time. The tool will accept a service name/URL as input and use automated testing, web scraping, API analysis, and potentially agentic capabilities to evaluate compliance with the 14-point service standard.

The assessment will be performed through real-time analysis rather than relying on static published datasets, providing an up-to-date evaluation of service compliance.

## Tasks (Ordered)

1. **Research and Document GDS Service Standard Points**
   - Document all 14 service standard points
   - Define measurable criteria for each point
   - Identify which aspects can be automated vs require manual review

2. **Define Assessment Architecture**
   - Design modular assessment engine
   - Define plugin/checker architecture for each standard point
   - Specify data collection methods (web scraping, API calls, accessibility testing)
   - Design scoring/reporting framework

3. **Implement Core Assessment Framework**
   - Build service discovery module (finding service URL from name)
   - Create assessment orchestrator
   - Implement result aggregation and reporting
   - Build CLI/web interface for initiating assessments

4. **Implement Standard-Specific Checkers**
   - Accessibility checker (WCAG compliance)
   - Security checker (HTTPS, headers, vulnerabilities)
   - Performance checker (load times, Core Web Vitals)
   - User research evidence detector
   - Open source/standards compliance checker
   - Service availability checker

5. **Develop Agentic Capabilities**
   - Implement intelligent navigation through service pages
   - Create content analysis for user needs assessment
   - Build pattern recognition for design patterns
   - Develop heuristic evaluation capabilities

6. **Create Reporting System**
   - Build detailed assessment reports
   - Implement scoring mechanism
   - Create recommendations engine
   - Export results in multiple formats (HTML, JSON, PDF)

7. **Testing and Validation**
   - Unit tests for each checker
   - Integration tests for full assessments
   - Validate against known compliant/non-compliant services
   - Performance and reliability testing

## Dependencies

### Technical Dependencies
- Web scraping library (Playwright/Puppeteer for JavaScript rendering)
- Accessibility testing tools (axe-core or pa11y)
- Security scanning tools (headers analysis, SSL/TLS checking)
- Performance testing tools (Lighthouse API)
- Natural language processing for content analysis
- API client libraries for government service APIs

### External Dependencies
- Access to government service URLs
- GDS Service Standard documentation (public)
- WCAG 2.1 AA guidelines
- Government design system patterns

### Sequencing Requirements
1. Core framework must be built before specific checkers
2. Service discovery needed before any assessment
3. Individual checkers can be developed in parallel
4. Reporting system depends on checker output format
5. Testing requires at least partial implementation of checkers

## Assumptions

1. **Service Accessibility**: Government services are publicly accessible via URLs
2. **Standard Stability**: GDS Service Standard points remain relatively stable during development
3. **Automation Feasibility**: At least 60-70% of standard points can be partially automated
4. **Public Information**: Services provide sufficient public-facing information for assessment
5. **Rate Limiting**: Government services allow reasonable automated access for testing
6. **Service Structure**: Services follow common patterns that can be programmatically analyzed

## Risks

### Technical Risks
- **Incomplete Automation**: Some standard points may be impossible to fully automate (e.g., "understand user needs")
- **False Positives/Negatives**: Automated checks may miss nuances or flag incorrect issues
- **Service Blocking**: Anti-bot measures might prevent automated assessment
- **Dynamic Content**: JavaScript-heavy sites may be difficult to analyze
- **API Limitations**: Not all services may have accessible APIs

### Operational Risks
- **Rate Limiting**: Frequent testing might trigger rate limits or IP blocks
- **Maintenance Burden**: Service changes might break checkers regularly
- **Legal Concerns**: Automated testing of government services might raise compliance issues
- **Resource Intensity**: Real-time comprehensive testing might be computationally expensive

### Quality Risks
- **Subjective Criteria**: Many service standard points involve subjective judgment
- **Context Loss**: Automated tools might miss important contextual information
- **Evolving Standards**: Service standards might change, requiring tool updates

## Out of Scope

1. **Manual Assessment Features**: Human-in-the-loop evaluation workflows
2. **Historical Tracking**: Storing and comparing assessment results over time
3. **Multi-jurisdiction Support**: Assessment of non-UK government service standards
4. **Remediation Tools**: Automated fixing of identified issues
5. **Authentication-Required Services**: Testing services behind login walls
6. **Cost Analysis**: Assessment of service delivery costs or value for money
7. **User Feedback Collection**: Direct surveying or feedback gathering from service users
8. **Compliance Certification**: Official certification or accreditation capabilities
9. **Service Integration**: Testing service-to-service integrations
10. **Load Testing**: Performance testing under high traffic conditions

## Open Questions

### Technical Questions
1. Which programming language/framework should be used? (Node.js/Python/Go?)
2. Should this be a CLI tool, web service, or both?
3. What level of caching is appropriate for assessment results?
4. How should we handle services requiring cookies/sessions?
5. Should we support headless vs headed browser testing?

### Scope Questions
1. Should the tool assess all 14 points or focus on technically measurable ones?
2. How deep should the assessment go (homepage only vs entire service)?
3. Should we include Welsh language service assessment?
4. Do we need to support assessment of mobile apps as well as web services?
5. Should the tool provide remediation guidance or just identify issues?

### Implementation Questions
1. What scoring methodology should be used (pass/fail, percentage, graded)?
2. How should partially-automated checks be presented?
3. What confidence levels should be assigned to automated findings?
4. Should the tool support batch assessment of multiple services?
5. How should we handle services with multiple entry points?

### Policy Questions
1. Are there any restrictions on automated testing of government services?
2. Should results be public or restricted to service teams?
3. How should the tool handle sensitive or security findings?
4. What disclaimers are needed about automated assessment limitations?
5. Should the tool integrate with existing GDS tools or platforms?