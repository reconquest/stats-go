package stats

import (
	"sync"
	"time"
)

type AverageDuration struct {
	*sync.Mutex
	name     string
	items    []time.Duration
	capacity int
}

func NewAverageDuration(capacity int) *AverageDuration {
	average := &AverageDuration{
		Mutex:    &sync.Mutex{},
		items:    []time.Duration{},
		capacity: capacity,
	}

	return average
}

func (average *AverageDuration) Push(duration time.Duration) {
	average.Lock()
	defer average.Unlock()

	average.items = append(average.items, duration)
	if average.capacity > 0 && len(average.items) > average.capacity {
		average.items = average.items[1:]
	}
}

func (average *AverageDuration) Get(count int) time.Duration {
	average.Lock()
	defer average.Unlock()

	if count == 0 {
		count = len(average.items)
	} else if len(average.items) < count {
		count = len(average.items)
	}

	if count == 0 {
		return 0
	}

	items := average.items[len(average.items)-count:]

	sum := time.Duration(0)
	for _, item := range items {
		sum += item
	}

	return time.Duration(sum / time.Duration(count))
}
