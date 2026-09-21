package LabOne

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Global variables shared between functions
var wg sync.WaitGroup

func addsAtomic(n int, total *atomic.Int64) bool {
	for i := 0; i < n; i++ {
		total.Add(1)
	}
	wg.Done() //lwt waitgroup know we have finished
	return true
}

func main() {
	var total atomic.Int64

	//for lopp using range option
	for range 10 {
		//the waitgroup is used as a barrier
		//init it to number of go routines
		wg.Add(1)
		go addsAtomic(1000, &total)
	}
	wg.Wait() // wait here until everyone (10 go routines) is done
	fmt.Println(total.Load())
}
