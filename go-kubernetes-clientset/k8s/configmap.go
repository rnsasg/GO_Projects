package k8s

import (
	"context"
	"log/slog"
	"time"

	common "github.com/rnsasg/GO_Projects/go-kubernetes-clientset"
	poll "github.com/rnsasg/GO_Projects/go-kubernetes-clientset/poll"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ConfigMapChecker is a type to adapt both kubenetes or fakeclient clientset
type ConfigMapChecker struct {
	clientset kubernetes.Interface
}

// NewConfigMapChecker creates a new instance of ConfigMapChecker
func NewConfigMapChecker(clientset kubernetes.Interface) *ConfigMapChecker {
	return &ConfigMapChecker{
		clientset: clientset,
	}
}

// IsConfigMapAvailable checks if ConfigMap is available within the given namespace
func (c *ConfigMapChecker) IsConfigMapAvailable(configMapName, namespace, modulename string, timeout time.Duration, maxRetries int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	config := poll.RetryConfig{
		InitialBackoff: 1 * time.Second,
		MaxBackoff:     5 * time.Minute,
		Timeout:        timeout,
		MaxRetries:     maxRetries,
		Mode:           poll.RetryTimeout,
	}
	slog.Info("IsConfigMapAvailable:", slog.String("Module", modulename), slog.Any("PollConfig", config))
	fn := func() bool {
		_, err := c.clientset.CoreV1().ConfigMaps(namespace).Get(ctx, configMapName, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				slog.Error("IsConfigMapAvailable:", slog.String("ConfigMap", configMapName), slog.String("not found in namespace", namespace), slog.String("error", err.Error()))
				return false
			}
			slog.Error("IsConfigMapAvailable:", slog.String("Error in reading configmap", configMapName), slog.String("error:", err.Error()), slog.String("error", err.Error()))
			return false
		}
		slog.Info("IsConfigMapAvailable:", slog.String("ConfigMap", configMapName), slog.String("found in namespace", namespace))
		return true
	}
	return poll.Retry(fn, config, modulename, "IsConfigMapAvailable")
}

// CreateConfigMap create a sample observability config map
func (c *ConfigMapChecker) CreateConfigMap(configMapName, namespace string, data map[string]string) (*corev1.ConfigMap, error) {
	cm := corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      configMapName,
			Namespace: namespace,
		},
		Data: data,
	}
	return c.clientset.CoreV1().ConfigMaps(namespace).Create(context.Background(), &cm, metav1.CreateOptions{})
}

func (c *ConfigMapChecker) GetConfigmap(name, namespace string) (*corev1.ConfigMap, error) {
	configmap, err := c.clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	return configmap, err
}

func (c *ConfigMapChecker) CreateObservabilityConfigmap(createAfter time.Duration) (*corev1.ConfigMap, error) {
	configData := map[string]string{
		"observability-config.json": `{
      		"cluster_cloud_account_id": "888888888888",
      		"cluster_cloud_provider": "AWS",
      		"cluster_name": "cluster",
      		"cluster_region": "us-west-2",
      		"http_ingestion_url": "https://sample.com",
      		"metrics_domain": "domain",
      		"metrics_storage_policy": "default",
      		"org_id": "orgid",
      		"environment": "staging"
   		}`,
	}
	<-time.After(createAfter)
	return c.CreateConfigMap(common.ObservabilityConfigMap, common.TanzuSystemNamespace, configData)
}
