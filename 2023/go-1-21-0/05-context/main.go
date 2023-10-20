package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	ctx := context.Background()

	// context.WithoutCancel

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Millisecond*10)
	defer cancel()

	ctxNoCancel := context.WithoutCancel(ctx)

	select {
	case <-ctxTimeout.Done():
		fmt.Println("ctxTimeout.Error:", ctxTimeout.Err())
		close(ch)
	}

	<-ch

	fmt.Println("context.WithoutCancel=Error:", ctxNoCancel.Err())

	// context.WithDeadlineCause (WithTimeoutCause works similarly)

	fmt.Println("")

	ch = make(chan struct{})
	ctx = context.Background()

	ctxDeadline, cancelDeadline := context.WithDeadlineCause(ctx, time.Now().Add(time.Millisecond*10), fmt.Errorf("deadline cause"))
	defer cancelDeadline()

	select {
	case <-ctxDeadline.Done():
		fmt.Println("ctxDeadline.Error:", ctxDeadline.Err())
		close(ch)
	}

	<-ch

	fmt.Println("context.WithDeadlineCause=Error:", ctxDeadline.Err())
	fmt.Println("context.WithDeadlineCause=Cause:", context.Cause(ctxDeadline))

	// context.AfterFunc

	fmt.Println("")

	ch = make(chan struct{})
	ctx = context.Background()

	ctxTimeout, cancelTimeout := context.WithTimeout(ctx, time.Millisecond*10)
	defer cancelTimeout()

	afterFunc := context.AfterFunc(ctxTimeout, func() {
		fmt.Println("AfterFunc called")
		close(ch)
	})

	select {
	case <-ctxTimeout.Done():
		fmt.Println("ctxTimeout.Error:", ctxTimeout.Err())
		afterFunc()
	}

	<-ch

	fmt.Println("context.WithTimeout=Error:", ctxDeadline.Err())
}
