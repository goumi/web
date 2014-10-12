package goumi

import "net/http"

// Context interface
type Context interface {

	// Readers for request and response
	Request() *http.Request
	Response() ResponseWriter

	// Iterator
	Next()

	// Add data to the chain
	Push(chain []Handler)
}

// Context is a wrapper f
type context struct {
	res ResponseWriter
	req *http.Request

	// Chain
	chain []Handler
}

// NewContext - Setup a new context
func newContext(w http.ResponseWriter, r *http.Request) Context {

	return &context{
		res: newResponse(w),
		req: r,
	}
}

// Getter fir Response
func (ctx *context) Response() ResponseWriter {
	return ctx.res
}

// Getter for Request
func (ctx *context) Request() *http.Request {
	return ctx.req
}

// Push - Append middleware to the context
func (ctx *context) Push(chain []Handler) {
	ctx.chain = append(chain, ctx.chain[0:]...)
}

// Next - Next function
func (ctx *context) Next() {

	// Check if we have middleware in the chain
	if len(ctx.chain) < 1 {
		return
	}

	// Grab the next middleware
	mw := ctx.chain[0]

	// Remove it from the chain
	ctx.chain = ctx.chain[1:]

	// Serve the middleware
	mw.Serve(ctx)
}
