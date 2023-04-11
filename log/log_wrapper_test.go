package log

import (
	"context"
	"github.com/smartpcr/azs-2-tf/utils/mocks"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var (
	settings = mocks.NewMockSettings()
)

func TestFileLogger(t *testing.T) {
	log := New(context.TODO(), settings, FileLogger)
	log.Info("Hello, world!")

	logFile := filepath.Join(settings.GetLogFolderPath(), settings.GetLogFileName())
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
	log := New(context.TODO(), settings, ConsoleLogger)
	log.Warnf("Hello, %s!", "World")
}
