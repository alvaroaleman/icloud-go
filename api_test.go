package icloudgo

import (
	"context"
	"os"
	"testing"
)

func TestNewFromCookieFile(t *testing.T) {
	file := os.Getenv("COOKIE_FILE")
	if file == "" {
		t.Skip("COOKIE_FILE not set")
	}
	if err := NewFromCookieFile(context.Background(), file); err != nil {
		t.Error(err)
	}
}
