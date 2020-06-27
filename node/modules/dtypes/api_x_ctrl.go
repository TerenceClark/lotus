package dtypes

import "sync"

type RPCHostVerifier struct {
	ValidHosts []string
	VHostsLock sync.RWMutex
}
