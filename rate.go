package stats

import (
	"sync"
	"time"
)

// Rate represents how many times counter was increased per given period of
// time.
type Rate struct {
	*sync.Mutex
	name     string
	interval time.Duration
	items    []time.Time
}

// NewRate creates new rate counter for given period of time.
func NewRate(interval time.Duration) *Rate {
	rate := &Rate{
		Mutex:    &sync.Mutex{},
		interval: interval,
		items:    []time.Time{},
	}

	go rate.sync()

	return rate
}

func (rate *Rate) sync() {
	for now := range time.Tick(rate.interval) {
		rate.tick(now.Add(rate.interval * 2 * -1))
	}
}

func (rate *Rate) tick(now time.Time) {
	rate.Lock()
	defer rate.Unlock()

	for i := 0; i < len(rate.items); i++ {
		item := rate.items[i]
		if item.After(now) {
			if i == 0 {
				return
			}

			rate.items = rate.items[i:]
			return
		}
	}
}

// Increase counter by 1.
func (rate *Rate) Increase() {
	rate.Lock()
	defer rate.Unlock()

	now := time.Now()
	rate.items = append(rate.items, now)
}

// Get returns number that describes how many times counter was increased per
// given time period.
func (rate *Rate) Get() int {
	rate.Lock()
	defer rate.Unlock()

	point := time.Now().Add(-rate.interval)

	for i := 0; i < len(rate.items); i++ {
		if rate.items[i].After(point) {
			return len(rate.items) - i
		}
	}

	return 0
}
