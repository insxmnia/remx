package slogger

import (
	"log/slog"
	"os"
	"sync"

	"github.com/google/uuid"
)

var (
	Instance   *slog.Logger
	once       sync.Once
	instanceId string
)

func Info(module, shortmsg string, args ...any) {
	fields := append([]any{"module", module, "instance_id", instanceId}, args...)
	slog.Info(shortmsg, fields...)
	if Instance != nil {
		Instance.Info(shortmsg, fields...)
	}
}

func Warn(module, shortmsg string, args ...any) {
	fields := append([]any{"module", module, "instance_id", instanceId}, args...)
	slog.Warn(shortmsg, fields...)
	if Instance != nil {
		Instance.Warn(shortmsg, fields...)
	}
}

func Error(module, shortmsg string, args ...any) {
	fields := append([]any{"module", module, "instance_id", instanceId}, args...)
	slog.Error(shortmsg, fields...)
	if Instance != nil {
		Instance.Error(shortmsg, fields...)
	}
}
func Fatal(module, shortmsg string, args ...any) {
	fields := append([]any{"module", module, "fatal", true, "instance_id", instanceId}, args...)
	slog.Error(shortmsg, fields...)
	if Instance != nil {
		Instance.Error(shortmsg, fields...)
	}
	os.Exit(1)
}

func init() {
	once.Do(func() {
		instanceId = uuid.NewString()
		filename := os.Getenv("LOG_OUTPUT_NAME")
		if filename == "" {
			filename = instanceId + ".log"
			slog.Warn("LOG_OUTPUT_NAME variable is missing, falling back to using instance ID", "filename", filename)
		}

		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}

		jsonHandler := slog.NewJSONHandler(file, nil)

		Instance = slog.New(jsonHandler)

	})
}
