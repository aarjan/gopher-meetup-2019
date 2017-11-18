package fib

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i <= n; i++ {
		if i == n {
			c <- x
		}
		x, y = y, x+y
	}
	close(c)
}

func Fib3(n int) int {
	c := make(chan int, n)
	go fibonacci(cap(c), c)
	return <-c
}
