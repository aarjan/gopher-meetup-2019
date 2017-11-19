/*
	Using 'time' package to get a ticker channel at sepecified period.
	Here the ticker, 't' recieves the input at every 1 second.
	Done channel is used to signal the end of ticking. 
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan int)
	t := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-t.C:
				fmt.Println("ticking")
			case <-done:
				fmt.Println("BOOM")
				break
			}
		}

	}()
	time.Sleep(3 * time.Second)
	done <- 0
}
