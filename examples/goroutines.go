package main

import "fmt"
import "time"

func printer(s string) {
	for i := 0; i <= 3; i++ {
		fmt.Println(s)
	}
}
func main() {
	go printer("goroutine")
	printer("direct")
	go printer("goroutine 2")
	time.Sleep(time.Millisecond)
}

/*
When we run this program, we see the output of the blocking call first,
then the interleaved output of the two gouroutines.
This interleaving reflects the goroutines being run concurrently by the Go runtime.
*/
