package main

import (
	"context"
	"fmt"
)

func main() {
	result, err := longRunningThingManager(context.Background(), "test")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func longRunningThingManager(ctx context.Context, data string) (string, error) {
	type wrapper struct {
		result string
		err    error
	}
	ch := make(chan wrapper, 1)
	go func() {
		// do the long running thing
		result, err := longRunningThingManager(ctx, data)
		ch <- wrapper{result, err}
	}()
	select {
	case data := <-ch:
		return data.result, data.err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
