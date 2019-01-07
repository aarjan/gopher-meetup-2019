package main

import "context"

type key string

var myKey = key("some-key")

func main() {
	// Use context Values only for request-scoped data that transits processes and
	// APIs, not for passing optional parameters to functions.
	valCtx := context.WithValue(context.Background(), myKey, "some-secret")

	// Canceling this context releases resources associated with it, so code should
	// call cancel as soon as the operations running in this Context complete:

	ctx, cancel := context.WithTimeout(valCtx, 2 * time.Second)
	defer cancel()

		result ,err := doSomething(ctx)	
	if err!=nil {
		panic(err)
	}
	fmt.Printf("Got result, %v", result)
}

func doSomething(ctx context.Context) (string, error) {
	c := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		c <- ctx.Value(myKey).(string)		
	}

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case val := <-c:
		return val, nil
	}
}