package stats

import (
	"sync"
)

var (
	stats = New()
)

type Stats struct {
	*sync.RWMutex
	items map[string]int
}

func New() *Stats {
	return &Stats{
		RWMutex: &sync.RWMutex{},
		items:   map[string]int{},
	}
}

func (stats *Stats) Increase(key string) {
	stats.Lock()
	defer stats.Unlock()

	stats.items[key] += 1
}

func (stats *Stats) IncreaseMany(key string, many int) {
	stats.Lock()
	defer stats.Unlock()

	stats.items[key] += many
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

func (stats *Stats) DecreaseMany(key string, many int) {
	stats.Lock()
	defer stats.Unlock()

	value, _ := stats.items[key]
	if value > 0 && value-many >= 0 {
		stats.items[key] = value - many
	} else {
		stats.items[key] = 0
	}
}

func (stats *Stats) Get() map[string]int {
	cloned := map[string]int{}

	stats.RLock()
	defer stats.RUnlock()

	for key, value := range stats.items {
		cloned[key] = value
	}

	return cloned
}

func Increase(key string) {
	stats.Increase(key)
}

func IncreaseMany(key string, many int) {
	stats.IncreaseMany(key, many)
}

func Decrease(key string) {
	stats.Decrease(key)
}

func DecreaseMany(key string, many int) {
	stats.DecreaseMany(key, many)
}

func Get() map[string]int {
	return stats.Get()
}
