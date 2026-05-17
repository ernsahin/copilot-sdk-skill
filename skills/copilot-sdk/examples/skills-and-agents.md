# Skills And Agents Pattern

Use this as a shape reference. Verify upstream Go SDK APIs before copying into production code.

Use SDK-loaded skills for reusable instructions and custom agents for scoped behavior.

```go
session, err := client.CreateSession(ctx, &copilot.SessionConfig{
	Model: "gpt-5",
	SkillDirectories: []string{
		"./skills",
	},
	CustomAgents: []copilot.CustomAgentConfig{
		{
			Name:        "code-reviewer",
			DisplayName: "Code Reviewer",
			Description: "Reviews repository changes with read-only tools and project review policy.",
			Tools:       []string{"grep", "glob", "view"},
			Prompt:      "Review code for correctness, maintainability, security, and failure behavior. Use repository search before conclusions.",
		},
	},
	Agent:               "code-reviewer",
	OnPermissionRequest: projectPermissionPolicy,
})
```

Keep prompts short. Put durable project rules in a skill. Put runtime permissions in permission handlers. Put external capabilities in tools or MCP servers.
