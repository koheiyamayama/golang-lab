package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
)

func main() {
	var errorSlices = []error{}
	errorSlices = append(errorSlices, errors.New("one error"))
	errorSlices = append(errorSlices, errors.New("two error"))

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})))

	var attrs []slog.Attr
	for i, err := range errorSlices {
		attrs = append(attrs, slog.String(fmt.Sprintf("%d", i), err.Error()))
	}

	gv := slog.GroupValue(attrs...)
	slog.ErrorContext(context.TODO(), "This is an error message", slog.Any("errors", gv))

	// {"time":"2024-02-29T16:04:13.9477+09:00","level":"ERROR","msg":"This is an error message","errors":{"0":"one error","1":"two error"}}
}
