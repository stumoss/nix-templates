// Package clilog contains utility types to help with configuring the logging
// via the CLI flags.
package clilog

import (
	"log/slog"
	"os"
)

var programLogLevel = new(slog.LevelVar)

type Level string

func (ll Level) AfterApply() error {
	switch ll {
	case Level(slog.LevelInfo.String()):
		programLogLevel.Set(slog.LevelInfo)
	case Level(slog.LevelWarn.String()):
		programLogLevel.Set(slog.LevelWarn)
	case Level(slog.LevelError.String()):
		programLogLevel.Set(slog.LevelError)
	case Level(slog.LevelDebug.String()):
		programLogLevel.Set(slog.LevelDebug)
	}

	return nil
}

type Mode string

func (lm Mode) AfterApply() error {
	var logger *slog.Logger
	switch lm {
	case "prod":
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: programLogLevel, AddSource: false}))
	case "dev":
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: programLogLevel, AddSource: true}))
	}

	slog.SetDefault(logger)

	return nil
}
