package logger

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name          string
		inputLevel    string
		expectedLevel slog.Level
	}{
		{"Debug Level", "debug", slog.LevelDebug},
		{"Info Level", "info", slog.LevelInfo},
		{"Warn Level", "warn", slog.LevelWarn},
		{"Error Level", "error", slog.LevelError},
		{"Invalid Level", "invalid", slog.LevelInfo},
		{"Upper Case Level", "DEBUG", slog.LevelDebug},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logBuffer := new(bytes.Buffer)
			levelVar := &slog.LevelVar{}
			handler := slog.NewJSONHandler(logBuffer, &slog.HandlerOptions{
				Level: levelVar,
			})
			logger := slog.New(handler)

			switch strings.ToLower(tc.inputLevel) {
			case "debug":
				levelVar.Set(slog.LevelDebug)
			case "info":
				levelVar.Set(slog.LevelInfo)
			case "warn":
				levelVar.Set(slog.LevelWarn)
			case "error":
				levelVar.Set(slog.LevelError)
			default:
				levelVar.Set(slog.LevelInfo)
			}

			logger.Debug("Debug message")
			logger.Info("Info message")
			logger.Warn("Warn message")
			logger.Error("Error message")

			decoder := json.NewDecoder(logBuffer)
			var foundLevels []string

			for {
				var logEntry map[string]interface{}
				if err := decoder.Decode(&logEntry); err != nil {
					break
				}

				if level, ok := logEntry["level"].(string); ok {
					foundLevels = append(foundLevels, level)
				}
			}

			expectedLevelStr := tc.expectedLevel.String()
			filteredLevels := []string{}
			for _, level := range foundLevels {
				if level == "INFO" || level == expectedLevelStr {
					filteredLevels = append(filteredLevels, level)
				}
			}

			expectedCount := 1
			if tc.expectedLevel == slog.LevelDebug {
				expectedCount = 2
			}

			if len(filteredLevels) != expectedCount {
				t.Errorf("Expected %d log entries at level %s, got %d",
					expectedCount, expectedLevelStr, len(filteredLevels))
			}
		})
	}
}
