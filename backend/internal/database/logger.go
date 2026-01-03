package database

import (
	"context"
	"encoding/json"
	"log"
	"time"
)

type Logger interface {
	Info(ctx context.Context, msg string, fields ...any)
	Error(ctx context.Context, msg string, err error, fields ...any)
}

type StdLogger struct {
	Service string
}

func (l StdLogger) Info(ctx context.Context, msg string, fields ...any) {
	l.log("info", msg, nil, fields...)
}

func (l StdLogger) Error(ctx context.Context, msg string, err error, fields ...any) {
	l.log("error", msg, err, fields...)
}

func (l StdLogger) log(
	level string,
	msg string,
	err error,
	fields ...any,
) {
	data := map[string]any{
		"level":   level,
		"service": l.Service,
		"message": msg,
		"time":    time.Now().UTC(),
	}

	if err != nil {
		data["error"] = err.Error()
	}

	if len(fields) > 0 {
		data["fields"] = fields
	}

	b, _ := json.Marshal(data)
	log.Println(string(b))
}
