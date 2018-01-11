package stats

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPercnetileDuration(t *testing.T) {
	test := assert.New(t)

	avg := NewPercentileDuration()
	avg.Push(time.Second * 4)
	avg.Push(time.Second * 3)
	avg.Push(time.Second * 2)
	avg.Push(time.Second * 2)
	avg.Push(time.Second * 2)
	avg.Push(time.Second * 2)
	avg.Push(time.Second * 2)
	avg.Push(time.Second * 2)
	avg.Push(time.Second * 2)
	avg.Push(time.Second * 2)

	test.Equal(time.Second*2, avg.Get(0, 30))
	test.Equal(time.Second*2, avg.Get(0, 80))
	test.Equal(time.Second*3, avg.Get(0, 85))
	test.Equal(time.Second*3, avg.Get(0, 90))
	test.Equal(time.Second*4, avg.Get(0, 100))
}
