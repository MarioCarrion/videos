package main

import (
	"context"
	"os"

	"golang.org/x/exp/slog"
)

type contextKey string

var authTokenContextKey contextKey = "AuthToken"

func NewAuthTokenContext(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, authTokenContextKey, value)
}

func AuthTokenFromContext(ctx context.Context) (string, bool) {
	value, exists := ctx.Value(authTokenContextKey).(string)

	return value, exists
}

type JSONHandler struct {
	slog.JSONHandler
}

func (h *JSONHandler) Handle(ctx context.Context, r slog.Record) error {
	if authToken, exists := AuthTokenFromContext(ctx); exists {
		r.AddAttrs(slog.String(string(authTokenContextKey), authToken))
	}

	return h.JSONHandler.Handle(ctx, r)
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout))

	// 1. Custom Handler
	// logger := slog.New(&JSONHandler{
	// 	JSONHandler: *slog.NewJSONHandler(os.Stdout),
	// })

	slog.SetDefault(logger)

	ctx := NewAuthTokenContext(context.Background(), "some fancy JWT")

	slog.InfoCtx(ctx, "info context")
}
