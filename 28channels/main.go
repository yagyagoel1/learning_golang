package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in golang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		val, IsChanelOpen := <-myCh
		fmt.Println(IsChanelOpen)
		fmt.Println(val)
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 5
		// myCh <- 6

		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
