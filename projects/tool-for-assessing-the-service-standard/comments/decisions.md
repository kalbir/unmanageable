# User Decisions for Planning Phase

## Technical Questions

1. **Programming language/framework**: Go
2. **Deployment model**: CLI tool  
3. **Result storage**: For iteration 1, no caching required but results should be written to file or database
4. **Credential handling**: Get as far as possible without credentials, future iterations may use dummy credentials
5. **Browser testing modes**: Both headless and headed browser capabilities needed

## Scope Questions

1. **Standard points to automate**: Focus on technically feasible ones for v1
2. **Service journey depth**: As far along the service as possible without credentials
3. **Manual assessment components**: No
4. **Service types**: Start with web services for v1
5. **Remediation**: Provide guidance on remediation

## Implementation Questions

1. **Scoring methodology**: Graded scoring
2. **Partial automation handling**: User could be prompted to check them, but have a mode/flag for including only fully automated vs. partially automated
3. **Assessment approach**: Neutral
4. **Batch processing**: No, one service at a time for now
5. **Multiple entry points**: Pick one entry point

## Policy Questions

1. **Legal restrictions**: User not aware of any
2. **Result visibility**: Public to whoever runs the tool but not published
3. **Accessibility findings**: Note them in the tool output
4. **Service notification**: None for this alpha version
5. **Government coordination**: No