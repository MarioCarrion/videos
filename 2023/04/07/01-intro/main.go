package main

import (
	"golang.org/x/exp/slog"
)

func main() {
	// 1. Default logger uses "log"
	// log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.LUTC)

	// 2. Text Handler
	// handler := slog.NewTextHandler(os.Stdout)
	// logger := slog.New(handler)
	// slog.SetDefault(logger)

	// 3. Enable lowel levels
	// handler := slog.HandlerOptions{
	// 	Level: slog.LevelDebug,
	// }.NewTextHandler(os.Stdout)
	//
	// logger := slog.New(handler)
	// slog.SetDefault(logger)

	//-

	slog.Debug("debugging ...", "something", "failed")
	slog.Info("info", "number", 10, "string", "some long \"sentence\"")

	// 4. Error when not using a string key: !BADKEY
	//slog.Warn("warning", 10, "ten") //, "something")
}
