package utils

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
	"time"
)

type LogHandler struct {
	writer    *os.File
	appName   string
	formatter func(level slog.Level, msg string, data map[string]interface{}) map[string]interface{}
}

// StructuredLogHandler returns a new LogHandler
func StructuredLogHandler(writer *os.File, appName string) *LogHandler {
	return &LogHandler{
		writer:  writer,
		appName: appName,
		formatter: func(level slog.Level, msg string, data map[string]interface{}) map[string]interface{} {
			return map[string]interface{}{
				"time":     time.Now().Format(time.RFC3339),
				"level":    level.String(),
				"message":  msg,
				"app_name": appName,
				"data":     data,
			}
		},
	}
}

// Enabled method to indicate which log levels are enabled
func (h *LogHandler) Enabled(_ context.Context, level slog.Level) bool {
	// Enable all log levels
	return true
}

// Handle implements the Handle method of the slog.Handler interface
func (h *LogHandler) Handle(ctx context.Context, record slog.Record) error {
	data := make(map[string]interface{})
	record.Attrs(func(attr slog.Attr) bool {
		data[attr.Key] = attr.Value.Any()
		return true
	})

	logData := h.formatter(record.Level, record.Message, data)
	logBytes, err := json.Marshal(logData)
	if err != nil {
		return err
	}
	_, err = h.writer.Write(logBytes)
	if err != nil {
		return err
	}
	_, err = h.writer.Write([]byte("\n"))
	return err
}

// WithAttrs implements the WithAttrs method of the slog.Handler interface
func (h *LogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	data := make(map[string]interface{})
	for _, attr := range attrs {
		data[attr.Key] = attr.Value.Any()
	}
	newFormatter := func(level slog.Level, msg string, existingData map[string]interface{}) map[string]interface{} {
		for key, value := range data {
			existingData[key] = value
		}
		return h.formatter(level, msg, existingData)
	}
	return &LogHandler{
		writer:    h.writer,
		appName:   h.appName,
		formatter: newFormatter,
	}
}

// WithGroup implements the WithGroup method of the slog.Handler interface
func (h *LogHandler) WithGroup(name string) slog.Handler {
	return h
}
