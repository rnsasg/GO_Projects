package k8s

import (
	"log/slog"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Clientset returns Kubernestes cluster clientset object to read the Kubernetes primitive metadata.
func Clientset() *kubernetes.Clientset {
	var (
		config *rest.Config
		err    error
	)

	config, err = rest.InClusterConfig()
	if err != nil {
		// development only
		localConfig := os.Getenv("KUBECONFIG")
		if localConfig == "" {
			slog.Error("Clientset: unable to obtain KUBECONFIG", "error", err.Error())
			os.Exit(1)
		}

		config, err = clientcmd.BuildConfigFromFlags("", localConfig)
		if err != nil {
			slog.Error("Clientset: failed to create k8s config", "err", err.Error())
			os.Exit(1)
		}
	}
	slog.Debug("Clientset: found InClusterConfig")

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		slog.Error("Clientset: failed to create", "error", err.Error())
		os.Exit(1)
	}

	return clientset
}
