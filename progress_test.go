package uiprogress_test

import (
	"github.com/gosuri/uiprogress"
	"time"
)

func ExampleStoppingPrintout() {
	uiprogress.Start()            // start rendering
	bar := uiprogress.AddBar(1) // Add a new bar
	bar.Incr()
	uiprogress.Stop()
	time.Sleep(1 * time.Second)
	// Output: blah

}