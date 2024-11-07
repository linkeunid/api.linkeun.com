package logger

import (
	"io"
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() *slog.Logger {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "logs/app.log",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   true,
	}

	multiWriter := io.MultiWriter(lumberjackLogger, os.Stdout)
	logger := slog.New(tint.NewHandler(multiWriter,
		&tint.Options{Level: slog.LevelDebug},
	))

	return logger
}
