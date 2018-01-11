package stats

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAverageDuration_CalculatesAverage(t *testing.T) {
	test := assert.New(t)

	avg := NewAverageDuration()
	avg.Push(time.Second)
	avg.Push(time.Second * 3)
	avg.Push(time.Second * 2)
	avg.Push(time.Second * 1)

	test.Equal(time.Second*2, avg.Get(3))
	test.Equal(time.Second*7/4, avg.Get(0))
	test.Equal(time.Second*7/4, avg.Get(4))
	test.Equal(time.Second*7/4, avg.Get(5))
	test.Equal(time.Second*3/2, avg.Get(2))
	test.Equal(time.Second*3/2, avg.Get(2))
}
