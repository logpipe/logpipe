package plugin

import (
	"sync"
)

var (
	inputLock      sync.RWMutex
	filterLock     sync.RWMutex
	outputLock     sync.RWMutex
	codecLock      sync.RWMutex
	inputBuilders  = make(map[string]InputBuilder)
	filterBuilders = make(map[string]FilterBuilder)
	outputBuilders = make(map[string]OutputBuilder)
	codecBuilders  = make(map[string]CodecBuilder)
)
