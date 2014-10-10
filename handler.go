package goumi

import "net/http"

// Handler - Serve a context
type Handler interface {
	Serve(ctx Context)
}

// HandlerFunc - Just build a handler
type HandlerFunc func(ctx Context)

// Serve calls f(ctx)
func (fn HandlerFunc) Serve(ctx Context) {
	fn(ctx)
}

// HTTP Handler Adapter
func HTTPHandler(h http.Handler) Handler {

	// Convert the HTTP Handler to a Goumi Hanler
	fn := func(ctx Context) {

		// Serve the handler
		h.ServeHTTP(ctx.Response(), ctx.Request())

		// Load the next context
		ctx.Next()
	}

	// Return the handler interface
	return HandlerFunc(fn)
}
