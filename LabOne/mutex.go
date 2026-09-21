package LabOne

import (
	"fmt"
	"sync"
)

// Global variables shared between functions -Bad idea
var wg sync.WaitGroup
var total int64

func adds(n int, theLock *sync.Mutex) bool {
	for i := 0; i < n; i++ {
		theLock.Lock()
		total++
		theLock.Unlock()
	}
	wg.Done() //let waitgroup know we have finished
	return true
}

func main() {
	//the lock will be passed by refrence between go routines
	//better than using a global variable
	var theLock sync.Mutex

	total = 0
	//the waitgroup is used as a barrier
	//init it to number of go routines
	wg.Add(10)

	//for loop using range option
	for i := range 10 {
		//staring
		fmt.Println(i)
		go adds(1000, &theLock)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done
	fmt.Println(total)
}
