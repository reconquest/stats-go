package stats

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRate(t *testing.T) {
	test := assert.New(t)

	rand.Seed(time.Now().Unix())

	rate := NewRate(time.Second)

	var count int
	var step int

	for range time.Tick(time.Second) {
		step++

		if step == 3 {
			test.Equal(count, rate.Get())
			return
		}

		count = rand.Intn(100)

		for i := 0; i < count; i++ {
			rate.Increase()
		}
	}

}
