# Skills And Agents Shape

This file is a shape reference only.

It is not source verification. Verify exact skill-loading and custom-agent APIs against the installed SDK or upstream source before writing final code.

## Conceptual Pattern

Use SDK-loaded skills for durable reusable instructions. Use custom agents for scoped behavior, tool access, and role-specific prompts.

Separate:

1. Skill directory configuration: where reusable instructions come from.
2. Custom agent definition: role, prompt, and allowed tools.
3. Active agent selection: which agent handles the session or task.
4. Permission policy: runtime allow, deny, or ask decisions.
5. Observability: events that show which skill or agent affected behavior.

## Quality Rules

1. Do not assume skills automatically apply to every custom agent.
2. Verify whether the target SDK needs explicit skill names on custom agents.
3. Keep custom agent prompts short.
4. Put durable rules in skills.
5. Put safety boundaries in permissions, tools, and host runtime code.
