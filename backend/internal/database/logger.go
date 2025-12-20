package database

import (
	"context"
	"log"
)

type Logger interface {
	Info(ctx context.Context, msg string, fields ...any)
	Error(ctx context.Context, msg string, fields ...any)
}

type StdLogger struct{}

func (l StdLogger) Info(ctx context.Context, msg string, fields ...any) {
	log.Println("[INFO]", msg, fields)
}

func (l StdLogger) Error(ctx context.Context, msg string, fields ...any) {
	log.Println("[ERROR]", msg, fields)
}
