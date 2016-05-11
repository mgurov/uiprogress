package uiprogress_test

import (
	"github.com/gosuri/uiprogress"
	"time"
	"os"
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