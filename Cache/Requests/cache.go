package requests

import (
	"sync"

	datasets "github.com/Chris-Kellett/workflow-manager/Datasets"
)

var (
	requests   map[string]datasets.RequestCache = make(map[string]datasets.RequestCache)
	requestsMu sync.RWMutex
)
