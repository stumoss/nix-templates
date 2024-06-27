package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/alecthomas/kong"

	"example.com/go-cli/internal/clilog"
)

var (
	_                  = kong.Must(&cli{})
	programName        = "<CHANGE ME>"
	programDescription = "<CHANGE ME>"
)

type cli struct {
	LogLevel clilog.Level `kong:"default='info',enum='error,warning,info,debug',help='logging level'"`
	LogMode  clilog.Mode  `kong:"default='prod',enum='prod,dev',help='logging mode'"`
}

func main() {
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer cancelCtx()

	cli := &cli{}
	k := kong.Parse(cli,
		kong.Name(programName),
		kong.Description("<CHANGE ME>"),
		kong.UsageOnError(),
		kong.DefaultEnvars(strings.ToUpper(programName)),
		kong.BindTo(ctx, (*context.Context)(nil)),
		kong.Bind(slog.Default()),
	)

	if err := k.Run(ctx, slog.Default()); err != nil {
		slog.Error("failed to run", slog.String("Error", err.Error()))
		os.Exit(1)
	}
}
