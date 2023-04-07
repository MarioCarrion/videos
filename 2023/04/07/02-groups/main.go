package main

import (
	"os"

	"golang.org/x/exp/slog"
)

func main() {
	handler := slog.NewTextHandler(os.Stdout)

	// 1. JSON Handler
	// handler := slog.NewJSONHandler(os.Stdout)

	//-

	group1 := slog.Group("app1", slog.Int("id", 1))
	group2 := slog.Group("app2", slog.Int("id", 2))

	logger := slog.New(handler)
	slog.SetDefault(logger)

	slog.Info("info", group1, group2)
}
