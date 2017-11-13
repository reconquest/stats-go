package stats

import (
	"sync"
)

type Stats struct {
	*sync.RWMutex
	items map[string]int64
}

func New() *Stats {
	return &Stats{
		RWMutex: &sync.RWMutex{},
		items:   map[string]int64{},
	}
}

func (stats *Stats) Increase(key string) {
	stats.Lock()
	defer stats.Unlock()

	stats.items[key] += 1
}

func (stats *Stats) Decrease(key string) {
	stats.Lock()
	defer stats.Unlock()

	value, _ := stats.items[key]
	if value > 0 {
		stats.items[key] = value - 1
	} else {
		stats.items[key] = 0
	}
}

func (stats *Stats) Get() map[string]int64 {
	cloned := map[string]int64{}

	stats.RLock()
	defer stats.RUnlock()

	for key, value := range stats.items {
		cloned[key] = value
	}

	return cloned
}
