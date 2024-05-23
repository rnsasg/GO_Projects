package poll

import (
	"log/slog"
	"time"
)

// RetryMode defines the mode of retry (Timeout or MaxRetries)
type RetryMode int

const (
	// RetryTimeout indicates retry based on maximum timeout
	RetryTimeout RetryMode = iota
	// RetryMaxRetries indicates retry based on maximum retries
	RetryMaxRetries
)

// RetryConfig contains configuration parameters for exponential backoff retry
type RetryConfig struct {
	InitialBackoff time.Duration
	MaxBackoff     time.Duration
	Timeout        time.Duration
	MaxRetries     int
	Mode           RetryMode
}

// String retrun Retry mode in string
func (r RetryMode) String() string {
	names := [...]string{
		"Timeout",
		"MaxRetries",
	}
	if r < RetryTimeout || r > RetryMaxRetries {
		return "Unknown"
	}
	return names[r]
}

// RetryFunc defines the function signature for retry/callback function
type RetryFunc func() bool

// Retry executes the provided function with retry logic based on the specified config
func Retry(fn RetryFunc, config RetryConfig, modulename, callbackFuncName string) bool {
	backoff := config.InitialBackoff
	retries := 0
	startTime := time.Now()
	slog.Info("Retry:", slog.String("Module", modulename), slog.String("Callback ", callbackFuncName), slog.Any(" Retry mode is", config.Mode.String()))

	if config.Mode.String() == "Unknown" {
		slog.Info("Retry:", slog.Any("Retry mode is invalid:", config.Mode))
		return false
	}
	for {
		if config.Mode == RetryTimeout && time.Since(startTime) >= config.Timeout {
			return false // Timeout reached, no need to retry further
		} else if config.Mode == RetryMaxRetries && retries >= config.MaxRetries {
			return false // Max retries reached, no need to retry further
		}

		if fn() {
			return true // Operation successful, no need to retry
		}

		// Exponential backoff
		slog.Info("Retry:", slog.String("Module", modulename), slog.String("Callback ", callbackFuncName), slog.Any(" will retry after", backoff))
		time.Sleep(backoff)
		backoff *= 2
		if backoff > config.MaxBackoff {
			slog.Info("Retry:", slog.String("Module", modulename), slog.String("Callback ", callbackFuncName), slog.Any(" resetting backoff", backoff))
			backoff = config.InitialBackoff
		}
		retries++
	}
}
