package main

import "context"

type WriteConnKey struct{}
type ReadConnKey struct{}
type SecKey struct{}

func WithWriteConn(ctx context.Context, writeConn int) context.Context {
	return context.WithValue(ctx, WriteConnKey{}, writeConn)
}

func WriteConnFromContext(ctx context.Context) int {
	v := ctx.Value(WriteConnKey{})
	if v == nil {
		return 1
	}
	return v.(int)
}

func WithReadConn(ctx context.Context, readConn int) context.Context {
	return context.WithValue(ctx, ReadConnKey{}, readConn)
}

func ReadConnFromContext(ctx context.Context) int {
	v := ctx.Value(ReadConnKey{})
	if v == nil {
		return 1
	}
	return v.(int)
}

func WithSec(ctx context.Context, sec int) context.Context {
	return context.WithValue(ctx, SecKey{}, sec)
}

func SecFromContext(ctx context.Context) int {
	v := ctx.Value(SecKey{})
	if v == nil {
		return 30
	}
	return v.(int)
}
