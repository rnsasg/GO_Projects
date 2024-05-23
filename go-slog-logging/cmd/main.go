package main

import "github.com/rnsasg/GO_Projects/go-slog-logging/api"

func main() {
	cfg := &api.RestAPIServerConfig{}
	// start rest for health probe and logging
	restServerConfig := api.NewRestAPIServer(cfg)
	restServerConfig.SetupDefaultLogLevel()
	go restServerConfig.Start()
}
