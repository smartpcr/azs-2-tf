package log

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFileLogger(t *testing.T) {
	log := New(context.TODO(), FileLogger)
	log.Info("Hello, world!")

	logFile := filepath.Join(logFolder, logFileName)
	_, err := os.Stat(logFile)
	if err != nil {
		t.Fatalf("log file doesn't exist: %v", err)
	}

	logContent, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	if string(logContent) == "" {
		t.Fatal("Log file is empty")
	}

	if !strings.Contains(string(logContent), "Hello, world!") {
		t.Fatal("Log file doesn't contain the expected message")
	}
}

func TestConsoleLogger(t *testing.T) {
	log := New(context.TODO(), ConsoleLogger)
	log.Warnf("Hello, %s!", "World")
}
