package stats

import (
	"sync"
	"time"
)

type AverageDuration struct {
	*sync.Mutex
	name  string
	items []time.Duration
}

func NewAverageDuration() *AverageDuration {
	average := &AverageDuration{
		Mutex: &sync.Mutex{},
		items: []time.Duration{},
	}

	return average
}

func (average *AverageDuration) Push(duration time.Duration) {
	average.Lock()
	defer average.Unlock()

	average.items = append(average.items, duration)
}

func (average *AverageDuration) Get(count int) time.Duration {
	average.Lock()

	if count == 0 {
		count = len(average.items)
	} else if len(average.items) < count {
		count = len(average.items)
	}

	items := average.items[len(average.items)-count:]

	average.Unlock()

	sum := time.Duration(0)
	for _, item := range items {
		sum += item
	}

	return time.Duration(sum / time.Duration(count))
}
