package io

import "context"

type Output interface {
	Print(i ...interface{})
	Println(i ...interface{})
	Printf(format string, i ...interface{})

	PrintErr(i ...interface{})
	PrintErrln(i ...interface{})
	PrintErrf(format string, i ...interface{})
}

var outputKey = struct{}{}

func WithOutput(ctx context.Context, out Output) context.Context {
	return context.WithValue(ctx, outputKey, out)
}

func OutputFrom(ctx context.Context) Output {
	return ctx.Value(outputKey).(Output)
}
