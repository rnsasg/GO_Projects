package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
	"unified-observability/config"

	"github.com/gorilla/mux"
)

// ErrRecoverable custom error type to handle retries
var (
	ErrInvalidLogLevel  = errors.New("invalid log level")
	ErrEmptyLogLevelEnv = errors.New("log level env is empty")
)

// RestAPICfg holds config for Rest API
// type RestAPICfg config.RestAPIServerConfig
// RestAPIServerConfig holds config for Rest API
type RestAPICfg struct {
	Port              int    `mapstructure:"port"`
	LogLevel          string `mapstructure:"log_level"`
	ReadHeaderTimeout time.Duration
	StopCh            chan struct{}
	Router            *mux.Router
	HTTPServer        *http.Server
}

// RestAPIServer a common interface for RestAPI
type RestAPIServer interface {
	Start()
	Stop()
	SetupDefaultLogLevel()
}

// SetupDefaultLogLevel sets loglevel during service mesh collector start up
func (rest *RestAPICfg) SetupDefaultLogLevel() {
	var err error
	if err = SetupLogLevelUsingEnv(); err == nil {
		slog.Info("Successfully set log level to %s using environment variable LOGLEVEL")
		return
	}

	if err = SetupLogLevelUsingConfigMap(rest.LogLevel); err == nil {
		slog.Info("SetupLogLevelUsingConfigMap", "Successfully set log level to %s using environment config map", rest.LogLevel)
		return
	}

	if err = SetLogLevel("info"); err != nil {
		slog.Info("Error in setting up loglevel to info")
		return
	}
	slog.Info("Successfully set default log level to info ")
}

// SetupLogLevelUsingConfigMap sets loglevel from configmap
func SetupLogLevelUsingConfigMap(loglevel string) error {

	slog.Info("SetupLogLevelUsingConfigMap", "Read config map log level value is", loglevel)
	if err := SetLogLevel(loglevel); err != nil {
		slog.Error("SetupLogLevelUsingEnv", "Error", err.Error())
		return err
	}
	return nil
}

// SetupLogLevelUsingEnv sets loglevel from environment variable
func SetupLogLevelUsingEnv() error {
	level, flag := os.LookupEnv("LOGLEVEL")
	if !flag {
		slog.Error("SetupLogLevelUsingEnv", "Error", ErrEmptyLogLevelEnv.Error())
		return ErrEmptyLogLevelEnv
	} else if err := SetLogLevel(level); err != nil {
		slog.Error("SetupLogLevelUsingEnv", "Error", err.Error())
		return err
	}
	return nil
}

// SetLogLevel sets specified log level
func SetLogLevel(level string) error {
	logLevel := &slog.LevelVar{}
	var logHandler *slog.TextHandler

	switch level {
	case "err":
		logLevel.Set(slog.LevelError)
	case "warn":
		logLevel.Set(slog.LevelWarn)
	case "info":
		logLevel.Set(slog.LevelInfo)
	case "debug":
		logLevel.Set(slog.LevelDebug)
	default:
		slog.Error("Invalid log level")
		return ErrInvalidLogLevel
	}
	logHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})
	slog.SetDefault(slog.New(logHandler))
	slog.Info("SetLogLevel", "Successfully set log level to %s", level)
	return nil
}

// Start starts the Rest API Server
func (rest *RestAPICfg) Start() {

	rest.Router.HandleFunc("/healthz", SetupReadyProbe).Methods(http.MethodGet)
	rest.Router.HandleFunc("/logging", SetupLoggingLevel).Methods(http.MethodGet)
	http.Handle("/", rest.Router)

	rest.HTTPServer.Addr = fmt.Sprintf(":%d", rest.Port)
	rest.HTTPServer.ReadHeaderTimeout = 30 * time.Second
	rest.HTTPServer.Handler = rest.Router

	if err := rest.HTTPServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("HTTP server error: %v", err)
	}
	log.Println("Stopped serving new connections.")
}

// Stop stop the Rest API Server
func (rest *RestAPICfg) Stop() {
	slog.Info("MUX RestServer : Shutting down rest server")
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := rest.HTTPServer.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete.")
}

// SetupReadyProbe serve heatlh request
func SetupReadyProbe(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// SetupLoggingLevel serve loggin request
func SetupLoggingLevel(w http.ResponseWriter, r *http.Request) {
	var level string
	level = r.URL.Query().Get("level")
	slog.Info("Logging", "Level ==> ", level)
	var message string = fmt.Sprintf("Successfully set loglevel to level ==> %s", level)

	if err := SetLogLevel(level); err != nil {
		message = err.Error()
	}
	w.Header().Set("Content-Type", "text/plain")

	bytes, err := w.Write([]byte(message))

	if err != nil {
		slog.Error("setupRestProbe : Error in returning response")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if bytes > 0 {
		slog.Debug("setupRestProbe", "Successfully wrote the %d", bytes)

	}
	w.WriteHeader(http.StatusOK)
}

// NewRestAPIServer returns a new RestAPIServerConfig object
func NewRestAPIServer(cfg *config.Config) RestAPIServer {
	return &RestAPICfg{
		Port: cfg.RestAPIServerConfig.Port,
		HTTPServer: &http.Server{
			ReadHeaderTimeout: cfg.RestAPIServerConfig.ReadHeaderTimeout * time.Second,
		},
		LogLevel: cfg.RestAPIServerConfig.LogLevel,
		Router:   mux.NewRouter(),
		StopCh:   make(chan struct{}, 1),
	}
}
