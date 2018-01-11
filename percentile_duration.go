package stats

import (
	"math"
	"sort"
	"sync"
	"time"
)

type PercentileDuration struct {
	*sync.Mutex
	name  string
	items []time.Duration
}

func NewPercentileDuration() *PercentileDuration {
	average := &PercentileDuration{
		Mutex: &sync.Mutex{},
		items: []time.Duration{},
	}

	return average
}

func (average *PercentileDuration) Push(duration time.Duration) {
	average.Lock()
	defer average.Unlock()

	average.items = append(average.items, duration)
}

func (average *PercentileDuration) Get(count, percents int) time.Duration {
	if percents == 0 {
		return 0
	}

	average.Lock()

	if count == 0 {
		count = len(average.items)
	} else if len(average.items) < count {
		count = len(average.items)
	}

	items := average.items[len(average.items)-count:]

	average.Unlock()

	if len(items) == 0 {
		return 0
	}

	sort.Slice(items, func(i, j int) bool { return items[i] < items[j] })

	top := int(math.Ceil(float64(len(items)*percents)/100)) - 1

	return items[top]
}
