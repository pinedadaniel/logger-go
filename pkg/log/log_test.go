package log

import (
	"context"
	"errors"
	"testing"
)

func TestInfoLogging(t *testing.T) {
	ctx := context.Background()
	Init(true)
	Info(ctx, "test info message",
		String("env", "test"),
		Int("code", 200),
	)
}

func TestErrorLogging(t *testing.T) {
	ctx := context.Background()
	testErr := errors.New("test error")
	Init(true)

	Error(ctx, "test error message",
		Err(testErr),
	)
}
