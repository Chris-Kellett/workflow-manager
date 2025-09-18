package requests

import datasets "github.com/Chris-Kellett/workflow-manager/Datasets"

func Get(correlationId string) (datasets.RequestCache, bool) {
	requestsMu.RLock()
	ret, ok := requests[correlationId]
	requestsMu.RUnlock()
	return ret, ok
}
