# Go Tool Pattern

Use this as a shape reference. Verify upstream Go SDK APIs before copying into production code.

Use SDK-native typed tools before hand-written orchestration.

```go
type LookupIssueParams struct {
	ID string `json:"id" jsonschema:"Issue identifier"`
}

lookupIssue := copilot.DefineTool(
	"lookup_issue",
	"Fetch issue details from the project tracker",
	func(params LookupIssueParams, inv copilot.ToolInvocation) (any, error) {
		if params.ID == "" {
			return nil, fmt.Errorf("issue id is required")
		}

		issue, err := fetchIssue(inv.TraceContext, params.ID)
		if err != nil {
			return nil, err
		}

		return map[string]any{
			"id":      issue.ID,
			"title":   issue.Title,
			"status":  issue.Status,
			"summary": issue.Summary,
		}, nil
	},
)

session, err := client.CreateSession(ctx, &copilot.SessionConfig{
	Model:               "gpt-5",
	Tools:               []copilot.Tool{lookupIssue},
	OnPermissionRequest: projectPermissionPolicy,
})
```

Review before production:

1. Are arguments validated?
2. Is the tool scope narrow?
3. Does the permission policy match the operation risk?
4. Are failures returned explicitly?
5. Does telemetry avoid leaking sensitive content?
