# logging

## how to use

```golang

package main

import (
	"github.com/WM1rr0rB8/librariesTest/backend/golang/tracing"
)

func main() {
	// logger from context
	logger := logging.L(ctx)
	// logger without access to context (level debug)
	logger = logging.Logger()

	// create logger with level
	logger = logging.LoggerWLevel("info")

	// add field to logger and put to context
	logger = logger.With(slog.String("key", "value"))
	ctx = logging.ContextWithLogger(ctx, l)
	
	// create logger with fields instantly
	logger = logging.WithAttrs(ctx, slog.String("key", "value"))
}

```
