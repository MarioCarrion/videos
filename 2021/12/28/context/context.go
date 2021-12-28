package main

import (
	"context"
)

func main() {
	ctx := context.Background()

	ParseIdiomatic(ctx, "hello")

	// XXX: Do not do this
	ParseNotIdiomatic("hello", ctx)

	// XXX: Do not do this
	_ = NotIdiomatic{
		ctx:   ctx,
		value: "something",
	}
}

func ParseIdiomatic(ctx context.Context, value string) {
}

func ParseNotIdiomatic(value string, ctx context.Context) {
}

type NotIdiomatic struct {
	ctx   context.Context
	value string
}
