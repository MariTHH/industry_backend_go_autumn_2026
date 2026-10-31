package main

import (
	"context"
	"errors"
	"sync"
)

type Result[R any] struct {
	Value R
	Err   error
}

var ErrInvalidWorkers = errors.New("workers must be positive")

func ParallelMap[T any, R any](ctx context.Context, workers int, in []T, fn func(context.Context, T) (R, error)) ([]Result[R], error) {
	if workers <= 0 {
		return nil, ErrInvalidWorkers
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	out := make([]Result[R], len(in))
	if len(in) == 0 {
		return out, nil
	}
	if workers > len(in) {
		workers = len(in)
	}
	jobs := make(chan int)
	var wg sync.WaitGroup
	for j := 0; j < workers; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				v, err := fn(ctx, in[i])
				if err != nil {
					var zero R
					out[i] = Result[R]{Value: zero, Err: err}
				} else {
					out[i] = Result[R]{Value: v}
				}
			}
		}()
	}
feed:
	for i := range in {
		select {
		case <-ctx.Done():
			break feed
		default:
		}
		select {
		case <-ctx.Done():
			break feed
		case jobs <- i:
		}
	}
	close(jobs)
	wg.Wait()
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return out, nil
}
