package uiprogress_test

import (
	"github.com/gosuri/uiprogress"
	"time"
	"testing"
	"sync"
)

func TestStoppingPrintout(t *testing.T) {

	counter := countingWriter{}

	progress := uiprogress.New()
	progress.RefreshInterval = time.Millisecond * 10
	progress.Out = &counter
	bar := progress.AddBar(1)
	progress.Start()
	bar.Incr()
	time.Sleep(time.Millisecond * 15)
	progress.Stop()
	time.Sleep(1 * time.Second)

	if 1 != counter.Count() {
		t.Errorf("Expected count 1, actual: %d", counter.Count())
	}
}

type countingWriter struct {
	count int
	mtx   sync.RWMutex
}

func (cw *countingWriter) Write(p []byte) (n int, err error) {
	cw.mtx.Lock()
	cw.count += 1
	cw.mtx.Unlock()
	return len(p), nil
}

func (cw *countingWriter) Count() int {
	cw.mtx.RLock()
	defer cw.mtx.RUnlock()
	return cw.count
}