package common

import (
	"math"
	"time"
)

// Observability common constants
const (
	ObservabilityConfigMap = "observability-config"
	TanzuSystemNamespace   = "tanzu-system"
	MaxTimeout             = time.Duration(math.MaxInt64)
	MaxRetries             = math.MaxInt64
	ObservabilityConfigKey = "observability-config.json"
)
