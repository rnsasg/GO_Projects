//go:build unit

package api_test

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"unified-observability/config"
	"unified-observability/pkg/api"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Conifg", func() {

	var (
		err           error
		restAPIServer api.RestAPIServer
		cfg           *config.Config
	)

	BeforeEach(func() {
		var err error
		err = os.Setenv("CONFIG_FILE_PATH", "../../test-run-config.json")
		Expect(err).ToNot(HaveOccurred())
		err = os.Setenv("TANZU_HUB_CONFIG_FILE_PATH", "../../test-observability-config.json")
		Expect(err).ToNot(HaveOccurred())

		cfg, err = config.Load(false)
		Expect(err).ToNot(HaveOccurred())

		restAPIServer = api.NewRestAPIServer(cfg)
	})

	It("should set log level to error", func() {
		err = os.Unsetenv("LOGLEVEL")
		Expect(err).ToNot(HaveOccurred())

		err = os.Setenv("LOGLEVEL", "err")
		Expect(err).ToNot(HaveOccurred())

		api.SetupLogLevelUsingEnv()
		Expect(true).To(Equal(slog.Default().Enabled(context.Background(), slog.LevelError)))
	})

	It("should set log level to warn", func() {
		err = os.Unsetenv("LOGLEVEL")
		Expect(err).ToNot(HaveOccurred())

		err = os.Setenv("LOGLEVEL", "warn")
		Expect(err).ToNot(HaveOccurred())

		api.SetupLogLevelUsingEnv()
		Expect(true).To(Equal(slog.Default().Enabled(context.Background(), slog.LevelWarn)))
	})

	It("should set log level to info", func() {
		err = os.Unsetenv("LOGLEVEL")
		Expect(err).ToNot(HaveOccurred())

		err = os.Setenv("LOGLEVEL", "info")
		Expect(err).ToNot(HaveOccurred())

		api.SetupLogLevelUsingEnv()

		Expect(true).To(Equal(slog.Default().Enabled(context.Background(), slog.LevelInfo)))
	})

	It("should set log level to debug", func() {
		err = os.Unsetenv("LOGLEVEL")
		Expect(err).ToNot(HaveOccurred())

		err = os.Setenv("LOGLEVEL", "debug")
		Expect(err).ToNot(HaveOccurred())
		api.SetupLogLevelUsingEnv()

		Expect(true).To(Equal(slog.Default().Enabled(context.Background(), slog.LevelDebug)))
	})

	It("should return error in case of invalid level", func() {
		err = os.Unsetenv("LOGLEVEL")
		Expect(err).ToNot(HaveOccurred())

		err := api.SetLogLevel("invalid")
		Expect(err).To(Equal(api.ErrInvalidLogLevel))
	})

	It("should return error in case of invalid level set through ENV", func() {
		err = os.Unsetenv("LOGLEVEL")
		Expect(err).ToNot(HaveOccurred())

		err = os.Setenv("LOGLEVEL", "invalid")
		Expect(err).ToNot(HaveOccurred())

		err := api.SetupLogLevelUsingEnv()
		Expect(err).To(Equal(api.ErrInvalidLogLevel))
	})

	It("should return error in case log level LOGLEVEL is not set", func() {
		err = os.Unsetenv("LOGLEVEL")
		Expect(err).ToNot(HaveOccurred())
		err := api.SetupLogLevelUsingEnv()
		Expect(err).To(Equal(api.ErrEmptyLogLevelEnv))
	})

	It("SetupDefaultLogLevel : should set the LOGLEVEL env", func() {
		err = os.Setenv("LOGLEVEL", "debug")
		Expect(err).ToNot(HaveOccurred())

		restAPIServer.SetupDefaultLogLevel()
		Expect(true).To(Equal(slog.Default().Enabled(context.Background(), slog.LevelDebug)))
	})

	It("SetupDefaultLogLevel : should set the LOGLEVEL to info if ENV and config map is not setting it", func() {
		err = os.Unsetenv("LOGLEVEL")
		Expect(err).ToNot(HaveOccurred())

		restAPIServer.SetupDefaultLogLevel()
		Expect(true).To(Equal(slog.Default().Enabled(context.Background(), slog.LevelInfo)))
	})

	It("SetupDefaultLogLevel : should set the LOGLEVEL to using configmap", func() {

		err = os.Unsetenv("LOGLEVEL")
		Expect(err).ToNot(HaveOccurred())

		err = api.SetLogLevel("info")
		Expect(err).ToNot(HaveOccurred())

		api.SetupLogLevelUsingConfigMap(cfg.RestAPIServerConfig.LogLevel)
		Expect(true).To(Equal(slog.Default().Enabled(context.Background(), slog.LevelDebug)))
	})

	It("SetupDefaultLogLevel : should set the LOGLEVEL to using rest api endpoint /logging", func() {

		// start rest for health probe and logging
		go restAPIServer.Start()
		defer restAPIServer.Stop()

		request, err := http.NewRequest("GET", "http://localhost:8080/logging?level=debug", nil)
		request.Header.Set("Content-Type", "text/plain; charset=utf-8")
		Expect(err).To(BeNil())

		// Send the request
		client := &http.Client{}
		response, err := client.Do(request)
		Expect(err).To(BeNil())
		Expect(response.StatusCode).To(Equal(http.StatusOK))
		defer response.Body.Close()
	})

})
