# Basic Go Session Shape

This file is a shape reference only.

It is not source verification. Do not claim exact SDK APIs are verified from this file. Do not copy exact imports, struct fields, event names, permission result names, or setup code from this file into a final answer.

Before implementation, verify the target SDK version from the installed package or upstream Go source.

## Conceptual Flow

1. Create or connect a Copilot SDK client.
2. Start the client if the target SDK requires explicit startup.
3. Create a session with model, working directory, tools, permissions, and event handling.
4. Register progress or streaming handlers before sending work.
5. Send the user task with a context that has timeout or cancellation.
6. Disconnect the session and stop owned runtime resources.

## Production Controls

Replace any approve-all prototype policy with scoped runtime permissions.

Define:

1. Authentication source and lifetime.
2. Working directory and config/state boundaries.
3. Permission behavior for read, write, shell, network, secrets, and external side effects.
4. Timeout and cancellation behavior.
5. Observable events and audit fields.
6. Cleanup behavior for normal completion, cancellation, and failure.
