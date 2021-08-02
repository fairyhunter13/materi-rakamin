package changoroutine

import (
	"runtime"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/require"
)

type counter struct {
	counter   int64
	goCounter uint64
	called    uint64
}

func (c *counter) AddCounter(inc int64) {
	atomic.AddUint64(&c.called, 1)
	pc := make([]uintptr, 1)
	count := runtime.Callers(5, pc)
	if count == 0 {
		atomic.AddUint64(&c.goCounter, 1)
	}
}

func TestTryChanGoroutine(t *testing.T) {
	// Don't modify this
	c := new(counter)
	intChan := make(chan int, 5)
	TryChanGoroutine(intChan, c)

	require.Equal(t, 5, len(intChan))
	allSum := 0
	for i := 1; i <= 3; i++ {
		select {
		case num := <-intChan:
			allSum += num
		default:
		}
	}

	require.Equal(t, 2, len(intChan))

	for i := 1; i <= 10; i++ {
		select {
		case num := <-intChan:
			allSum += num
		default:
		}
	}

	require.Equal(t, 0, len(intChan))
	require.Equal(t, 47, allSum)

	require.EqualValues(t, 50, c.called)
	require.True(t, c.counter >= 10*100*10)
	require.True(t, c.goCounter >= 45)
}
