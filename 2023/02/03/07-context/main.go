package main

import (
	"context"
	"errors"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	fmt.Printf("error is: %v\n", ctx.Err())
	fmt.Printf("cause was: %v\n", context.Cause(ctx))

	//-

	ctxCause, cancelCause := context.WithCancelCause(context.Background())
	cancelCause(errors.New("something failed"))

	fmt.Printf("error is: %v\n", ctxCause.Err())
	fmt.Printf("cause was: %v\n", context.Cause(ctxCause))
}
