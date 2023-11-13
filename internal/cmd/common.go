package cmd

import (
	"fmt"

	"github.com/merlindorin/hoomy/internal/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Commons struct {
	Development bool   `short:"D" env:"DEBUG,DEV,DEVELOPMENT" help:"Enable development mode"`
	Level       string `short:"l" env:"LOG_LEVEL" help:"Set the logging level (debug|info|warn|error|fatal)" default:"info"`

	Version Version `cmd:"" help:"Print version information and quit."`
	Licence Licence `cmd:"" help:"Print licence of the application."`
}

func (g *Commons) Logger() (*zap.Logger, error) {
	level, err := zapcore.ParseLevel(g.Level)
	if err != nil {
		return nil, fmt.Errorf("cannot parse logger level \"%s\": %w", g.Level, err)
	}

	if g.Development {
		level = zapcore.DebugLevel
	}

	return logger.New(level, g.Development)
}
