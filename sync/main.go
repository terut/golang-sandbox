package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(sleep time.Duration) {
			defer wg.Done()

			fmt.Println("sleep: ", sleep*time.Second)
			time.Sleep(sleep * time.Second)
		}(time.Duration(i + 1))
	}

	wg.Wait()
	fmt.Println("exit wait group")
}
