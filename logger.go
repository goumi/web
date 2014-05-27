package ape

import (
	"log"
	"net/http"
	"os"
	"time"
)

// Logger is a middleware handler that logs the request as it goes in and the response as it goes out.
type Logger struct {
	// Logger is the log.Logger instance used to log messages with the Logger middleware
	*log.Logger
}

// NewLogger returns a new Logger instance
func NewLogger() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", 0),
	}
}

func (l *Logger) Serve(ctx *Context) {
	start := time.Now()
	l.Printf("-> %s %s", ctx.Req.Method, ctx.Req.URL.Path)

	// Load the next middleware
	ctx.Next()

	// Load the body length
	size := ByteSize(ctx.Res.Body.Len()).String()

	// Show the response
	l.Printf("<- %v %s - %v %v", ctx.Res.Code, http.StatusText(ctx.Res.Code), time.Since(start), size)
}
