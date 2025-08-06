package logger

import (
	"io"
	"log/slog"
	"os"
)

type Config struct {
	AppENV  string
	LogPath string
	Debug   bool
}

func SetupLogger(config *Config) {
	var writer io.Writer
	var logLevel slog.Level

	// Determine log level
	if config.AppENV == "production" {
		logLevel = slog.LevelWarn
	} else {
		if config.Debug {
			logLevel = slog.LevelDebug
		} else {
			logLevel = slog.LevelInfo
		}
	}

	if config.LogPath == "" {
		// If no log path is specified, log only to stdout
		writer = os.Stdout
	} else {
		// Ensure log directory exists
		err := os.MkdirAll(config.LogPath, 0755)
		if err != nil {
			panic(err)
		}

		// Open log file
		file, err := os.OpenFile(config.LogPath+"/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}

		// Write to both stdout and file
		writer = io.MultiWriter(os.Stdout, file)
	}

	logJSONHandler := slog.NewJSONHandler(writer, &slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
	})
	logger := slog.New(logJSONHandler)
	slog.SetDefault(logger)

	if config.LogPath == "" {
		slog.Info("logging only to stdout because LogPath is empty")
	} else {
		if config.AppENV == "production" {
			slog.Info("logging to file and stdout")
		} else {
			slog.Info("logging to stdout and file (non-production mode)")
		}
	}
}
