package LabOne

import (
	"fmt"
	"sync"
	"time"
)

//Gloabal variables shared between functions -Bad idea

func main() {
	var wg sync.WaitGroup
	barrier := make(chan bool)

	doStuffOne := func() bool {
		fmt.Println("StuffOne - Part A")
		//wait here
		fmt.Println("StuffOne - Part B")
		wg.Done()
		return true
	}
	doSuffTwo := func() bool {
		time.Sleep(time.Second * 5)
		fmt.Println("StuffTwo - Part A")
		//wait here
		<-barrier
		fmt.Println("StuffTwo - Part B")
		wg.Done()
		return true
	}
	wg.Add(2)
	go doStuffOne()
	go doSuffTwo()
	wg.Wait() //wait here untill everyone (10 go routines) is done
}
