package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("race condition ")
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}
	var score = []int{0}
	wg.Add(3)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("One R")
		m.Lock()
		score = append(score, 1)
		m.Unlock()

	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("two R")
		m.Lock()
		score = append(score, 2)
		m.Unlock()

	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("three R")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("three R")
		m.RLock()
		fmt.Println(score)
		m.RUnlock()
	}(wg, mut)
	wg.Wait()
	defer fmt.Println("", score)
}
