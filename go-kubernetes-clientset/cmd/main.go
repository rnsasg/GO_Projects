package main

import (
	"log/slog"
	"math"
	"os"
	"time"

	commonK8s "github.com/rnsasg/GO_Projects/go-kubernetes-clientset/k8s"
)

const (
	ObservabilityConfigMap = "observability-config"
	TanzuSystemNamespace   = "tanzu-system"
	MaxTimeout             = time.Duration(math.MaxInt64)
	MaxRetries             = math.MaxInt64
	ObservabilityConfigKey = "observability-config.json"
)

func main() {
	clientset := commonK8s.Clientset()
	checker := commonK8s.NewConfigMapChecker(clientset)
	ok := checker.IsConfigMapAvailable(ObservabilityConfigMap, TanzuSystemNamespace, "Service-Mesh-Collector", MaxTimeout, MaxRetries)
	if !ok {
		slog.Error("Could not find observability configmap")
		os.Exit(1)
	}
}
