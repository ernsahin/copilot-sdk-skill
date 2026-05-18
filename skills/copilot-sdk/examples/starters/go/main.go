package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	copilot "github.com/github/copilot-sdk/go"
)

func main() {
	prompt := "Say hello in one short sentence."
	if len(os.Args) > 1 {
		prompt = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	client := copilot.NewClient(&copilot.ClientOptions{LogLevel: "error"})
	if err := client.Start(ctx); err != nil {
		log.Fatalf("start copilot client: %v", err)
	}
	defer client.Stop()

	done := make(chan struct{})
	session, err := client.CreateSession(ctx, &copilot.SessionConfig{
		Model:               "gpt-5",
		Streaming:           true,
		OnPermissionRequest: conservativePermission,
	})
	if err != nil {
		log.Fatalf("create session: %v", err)
	}
	defer session.Disconnect()

	session.On(func(event copilot.SessionEvent) {
		switch data := event.Data.(type) {
		case *copilot.AssistantMessageDeltaData:
			fmt.Print(data.DeltaContent)
		case *copilot.AssistantMessageData:
			fmt.Print(data.Content)
		case *copilot.SessionIdleData:
			close(done)
		}
	})

	if _, err := session.Send(ctx, copilot.MessageOptions{Prompt: prompt}); err != nil {
		log.Fatalf("send prompt: %v", err)
	}

	select {
	case <-done:
		fmt.Println()
	case <-ctx.Done():
		log.Fatalf("wait for response: %v", ctx.Err())
	}
}

func conservativePermission(
	request copilot.PermissionRequest,
	invocation copilot.PermissionInvocation,
) (copilot.PermissionRequestResult, error) {
	switch request.Kind {
	case copilot.PermissionRequestKindRead:
		return copilot.PermissionRequestResult{Kind: copilot.PermissionRequestResultKindApproved}, nil
	default:
		fmt.Fprintf(os.Stderr, "denied %s request in session %s\n", request.Kind, invocation.SessionID)
		return copilot.PermissionRequestResult{Kind: copilot.PermissionRequestResultKindUserNotAvailable}, nil
	}
}
