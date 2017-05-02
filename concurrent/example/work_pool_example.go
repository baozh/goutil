package main

import (
	"fmt"
	"sync"

	"github.com/baozh/goutil/concurrent"
)

// names provides a set of names to display.
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

var wg sync.WaitGroup

// namePrinter provides special support for printing names.
type namePrinter struct {
	name string
}

// Task implements the Worker interface.
func (m *namePrinter) print() {
	fmt.Println(m.name)
	wg.Done()
}

// main is the entry point for all Go programs.
func main() {
	// Create a work pool with 5 goroutines.
	pool := concurrent.NewWorkerPool(5, 50)

	wg.Add(100 * len(names))
	for i := 0; i < 100; i++ {
		// Iterate over the slice of names.
		for _, name := range names {
			// Create a namePrinter and provide the
			// specific name.
			tmp := namePrinter{
				name: name,
			}

			fn := func() {
				// Submit the task to be worked on. When RunTask
				// returns we know it is being handled.
				tmp.print()
			}
			pool.Run(fn)
		}
	}
	wg.Wait()

	// Shutdown the work pool and wait for all existing work
	// to be completed.
	pool.ShutDown()
}
