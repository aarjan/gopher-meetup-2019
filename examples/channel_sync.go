/*
	Select statement can be used to work with multiple channels.
	Unbuffered channel have a blocking nature of execution, but Select statment can be used to switch between fastest input
	while working with multiple channels.
*/

package main

import "fmt"

func run(ch chan int, done chan int) {
	for {
		select {
		case <-ch:
			fmt.Println("getting")
		case <-done:
			fmt.Println("exiting")
		default:
			fmt.Println("waiting")
		}
	}
}
func main() {
	ch, done := make(chan int), make(chan int)
	go run(ch, done)
	for i := 0; i <= 10; i++ {
		ch <- i
	}
	done <- 1

}
