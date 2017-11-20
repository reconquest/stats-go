package stats

import (
	"sync"
	"time"
)

type Rate struct {
	*sync.Mutex
	name     string
	interval time.Duration
	items    []time.Time
}

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

func (rate *Rate) Increase() {
	rate.Lock()
	defer rate.Unlock()

	now := time.Now()
	rate.items = append(rate.items, now)
}

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
