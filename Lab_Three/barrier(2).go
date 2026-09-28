package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, total int, arrived *int,
	theLock *sync.Mutex, sem *semaphore.Weighted, ctx context.Context) bool {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	//barrier starts
	theLock.Lock()
	*arrived++
	if *arrived == total {
		sem.Release(1)
	}
	theLock.Unlock()
	sem.Acquire(ctx, 1)
	sem.Release(1)
	//barrier end
	fmt.Println("PartB", goNum)
	return true
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	ctx := context.TODO()
	var theLock sync.Mutex
	var arrived = 0
	sem := semaphore.NewWeighted(int64(totalRoutines))
	sem.Acquire(ctx, 1)
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &wg, totalRoutines, &arrived, &theLock, sem, ctx)
	}

	wg.Wait() //wait for everyone to finish before exiting
}
