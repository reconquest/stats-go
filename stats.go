package stats

import (
	"sync"
)

var (
	stats = NewStats()
)

// Stats represents numeric statistics. Simplest counter that can be increased
// and decreased. Thread-safe.
type Stats struct {
	*sync.RWMutex
	items map[string]int
}

// NewStats creates new stats. Use this method if global stats is not enough
// for you.
func NewStats() *Stats {
	return &Stats{
		RWMutex: &sync.RWMutex{},
		items:   map[string]int{},
	}
}

// Set given value for given key.
func (stats *Stats) Set(key string, value int) {
	stats.Lock()
	defer stats.Unlock()

	stats.items[key] = value
}

// Increase counter with given key by 1.
func (stats *Stats) Increase(key string) {
	stats.Lock()
	defer stats.Unlock()

	stats.items[key] += 1
}

// IncreaseMany counter with given key by given `many` number.
func (stats *Stats) IncreaseMany(key string, many int) {
	stats.Lock()
	defer stats.Unlock()

	stats.items[key] += many
}

// Decrease counter with  key by 1, value will not go less than zero.
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

// DecreaseMany counter with key by given `many` number, value will not go less
// than zero.
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

// Get returns table with specified keys and counters.
func (stats *Stats) Get() map[string]int {
	cloned := map[string]int{}

	stats.RLock()
	defer stats.RUnlock()

	for key, value := range stats.items {
		cloned[key] = value
	}

	return cloned
}

// Set given value for given key.
func Set(key string, value int) {
	stats.Set(key, value)
}

// Increase counter with given key by 1.
func Increase(key string) {
	stats.Increase(key)
}

// IncreaseMany counter with given key by given `many` number.
func IncreaseMany(key string, many int) {
	stats.IncreaseMany(key, many)
}

// Decrease counter with  key by 1, value will not go less than zero.
func Decrease(key string) {
	stats.Decrease(key)
}

// DecreaseMany counter with key by given `many` number, value will not go less
// than zero.
func DecreaseMany(key string, many int) {
	stats.DecreaseMany(key, many)
}

// Get returns table with specified keys and counters.
func Get() map[string]int {
	return stats.Get()
}
