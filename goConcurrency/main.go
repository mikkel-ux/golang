package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	println("CPU Cores: ", runtime.NumCPU())

	counter := 0
	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	println("Counter might be wrong:", counter)

	println("\n--- WITH Mutex (Protected) ---")
	counter = 0
	var mu sync.Mutex

	wg = sync.WaitGroup{}
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Counter is correct:", counter)
}
