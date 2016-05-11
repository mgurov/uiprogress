package uiprogress_test

import (
	"github.com/gosuri/uiprogress"
	"time"
	"os"
	"testing"
)

func ExampleStoppingPrintout() {
	progress := uiprogress.New()            // start rendering
	progress.RefreshInterval = time.Millisecond * 10
	progress.Out = os.Stdout
	progress.Start()
	bar := progress.AddBar(1) // Add a new bar
	bar.Incr()
	time.Sleep(time.Millisecond * 15)
	//progress.Bars = nil
	progress.Stop()
	time.Sleep(1 * time.Second)
	// Output: [====================================================================]

}

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

	if 1 != counter.Count {
		t.Errorf("Expected count 1, actual: %d", counter.Count)
	}
}

type countingWriter struct {
	Count int
}

func (cw *countingWriter) Write(p []byte) (n int, err error) {
	cw.Count ++
	return len(p), nil
}
