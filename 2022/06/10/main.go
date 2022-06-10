package main

import (
	"context"
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
)

type Service func(ctx context.Context) (int64, error)

func main() {
	svcs := []Service{
		Service1,
		Service2,
		Service3,
	}

	var total int64

	start := time.Now()

	g, ctx := errgroup.WithContext(context.Background())

	for _, svc := range svcs {
		svc := svc
		g.Go(func() error {
			v, err := svc(ctx)
			if err != nil {
				return err
			}

			_ = atomic.AddInt64(&total, v)

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatalln("Group", err)
	}

	fmt.Println("Total", total, "Duration", time.Now().Sub(start))
}

func Service1(ctx context.Context) (int64, error) {
	time.Sleep(100 * time.Millisecond)
	return 1, nil
}

func Service2(ctx context.Context) (int64, error) {
	time.Sleep(100 * time.Millisecond)
	return 2, nil
}

func Service3(ctx context.Context) (int64, error) {
	time.Sleep(200 * time.Millisecond)
	return 3, nil
}
