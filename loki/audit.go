package main

import (
	"log/slog"
	"math/rand"
	"os"
	"time"
)

type LogEntry struct {
	UseCase      string
	UseCaseType  string
	TraceID      int
	SpanID       int
	ParentSpanID int
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// List of use cases and types
	useCases := []string{"Authentication", "PaymentProcessing", "OrderManagement", "Analytics"}
	useCaseTypes := []string{"Read", "Write", "Update", "Delete"}

	// Define the log directory and file
	logDir := "/app/logs"
	logFile := logDir + "/app.log"

	// Check if the directory exists
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll(logDir, 0755)
	}

	// Open the existing log file
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}
	defer file.Close()

	// Create a logger that writes to the log file
	logger := slog.New(slog.NewTextHandler(file, nil))

	// Continuous loop to generate logs
	for {
		// Generate random values
		traceID := rand.Intn(1000) + 1
		spanID := rand.Intn(1000) + 1
		parentSpanID := rand.Intn(1000) + 1
		useCase := useCases[rand.Intn(len(useCases))]
		useCaseType := useCaseTypes[rand.Intn(len(useCaseTypes))]

		// Create a log entry
		entry := LogEntry{
			UseCase:      useCase,
			UseCaseType:  useCaseType,
			TraceID:      traceID,
			SpanID:       spanID,
			ParentSpanID: parentSpanID,
		}

		// Log the entry
		logger.Info("Log Entry",
			slog.String("use_case", entry.UseCase),
			slog.String("use_case_type", entry.UseCaseType),
			slog.Int("trace_id", entry.TraceID),
			slog.Int("span_id", entry.SpanID),
			slog.Int("parent_span_id", entry.ParentSpanID),
		)

		// Sleep to control log generation frequency
		time.Sleep(1 * time.Second)
	}
}
