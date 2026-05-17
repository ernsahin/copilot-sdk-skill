# Basic Go Session

Use this as a shape reference, not as a frozen API guarantee. Verify upstream before final code.

```go
package main

import (
	"context"
	"fmt"
	"log"

	copilot "github.com/github/copilot-sdk/go"
)

func main() {
	ctx := context.Background()

	client := copilot.NewClient(&copilot.ClientOptions{
		LogLevel: "error",
	})
	if err := client.Start(ctx); err != nil {
		log.Fatal(err)
	}
	defer client.Stop()

	session, err := client.CreateSession(ctx, &copilot.SessionConfig{
		Model:               "gpt-5",
		OnPermissionRequest: copilot.PermissionHandler.ApproveAll,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer session.Disconnect()

	done := make(chan struct{})
	session.On(func(event copilot.SessionEvent) {
		switch d := event.Data.(type) {
		case *copilot.AssistantMessageData:
			fmt.Println(d.Content)
		case *copilot.SessionIdleData:
			close(done)
		}
	})

	if _, err := session.Send(ctx, copilot.MessageOptions{
		Prompt: "Summarize this repository's test strategy.",
	}); err != nil {
		log.Fatal(err)
	}

	<-done
}
```

For production, replace approve-all with a scoped permission policy and add cancellation, timeout, auth, and observability decisions.

