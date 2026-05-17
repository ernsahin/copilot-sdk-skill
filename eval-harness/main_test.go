package main

import (
	"os"
	"testing"

	copilot "github.com/github/copilot-sdk/go"
)

func TestReadOnlyPermissionHandlerApprovesRead(t *testing.T) {
	file := tempLog(t)
	path := "README.md"
	result, err := readOnlyPermissionHandler(file)(copilot.PermissionRequest{
		Kind: copilot.PermissionRequestKindRead,
		Path: &path,
	}, copilot.PermissionInvocation{SessionID: "s1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Kind != copilot.PermissionRequestResultKindApproved {
		t.Fatalf("expected read approval, got %s", result.Kind)
	}
}

func TestReadOnlyPermissionHandlerDeniesWrite(t *testing.T) {
	file := tempLog(t)
	path := "README.md"
	result, err := readOnlyPermissionHandler(file)(copilot.PermissionRequest{
		Kind:     copilot.PermissionRequestKindWrite,
		FileName: &path,
	}, copilot.PermissionInvocation{SessionID: "s1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Kind != copilot.PermissionRequestResultKindUserNotAvailable {
		t.Fatalf("expected write denial, got %s", result.Kind)
	}
}

func tempLog(t *testing.T) *os.File {
	t.Helper()
	file, err := os.CreateTemp(t.TempDir(), "permissions-*.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = file.Close()
	})
	return file
}
